package database

import (
	"encoding/json"
	"io/ioutil"
)

type Genesis struct {
	Balances map[Account]uint `json:"balances""`
	Symbol   string           `json:"symbol"`
}

func loadGenesis(path string) (Genesis, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var loadedGenesis Genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}
