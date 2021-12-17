package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type State struct {
	Balances   map[Account]uint
	trxMempool []Transaction
	dbFile     *os.File
}

func (s *State) apply(trx Transaction) error {
	if trx.IsReward() {
		s.Balances[trx.To] += trx.Value
		return nil
	}

	if trx.Value > s.Balances[trx.From] {
		return fmt.Errorf("insufficient balance")
	}

	s.Balances[trx.From] -= trx.Value
	s.Balances[trx.To] += trx.Value

	return nil
}

// Add Insert transaction to mempool
func (s *State) Add(trx Transaction) error {
	if err := s.apply(trx); err != nil {
		return err
	}

	s.trxMempool = append(s.trxMempool, trx)

	return nil
}

// Persist store transactions data to state's trxMempool
func (s *State) Persist() error {
	mempool := make([]Transaction, len(s.trxMempool))
	copy(mempool, s.trxMempool)

	for i := 0; i < len(mempool); i++ {
		trxJson, err := json.Marshal(mempool[i])
		if err != nil {
			return err
		}

		if _, err = s.dbFile.Write(append(trxJson, '\n')); err != nil {
			return err
		}

		s.trxMempool = s.trxMempool[1:]
	}

	return nil
}

// Close safely closes dbFile from reading
func (s *State) Close() {
	s.dbFile.Close()
}

// NewStateFromDisk loads state from disk
func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	genFilePath := filepath.Join(cwd, "internal/database", "genesis.json")
	gen, err := loadGenesis(genFilePath)
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	txDbFilePath := filepath.Join(cwd, "internal/database", "trx.db")
	f, err := os.OpenFile(txDbFilePath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	state := &State{balances, make([]Transaction, 0), f}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var tx Transaction
		json.Unmarshal(scanner.Bytes(), &tx)

		if err := state.apply(tx); err != nil {
			return nil, err
		}
	}

	return state, nil
}
