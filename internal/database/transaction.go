package database

type Transaction struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
}

// NewTransaction creates a Transaction
func NewTransaction(from Account, to Account, value uint, data string) Transaction {
	return Transaction{from, to, value, data}
}

// IsReward check whether transaction is a reward
func (t *Transaction) IsReward() bool {
	return t.Data == "reward"
}
