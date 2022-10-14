// Package TRansaction represents the events that make changes to the blockchain and accounts
package transaction

import "time"

// Transaction is the fundamental entry for recording changes to the blockchain
type Transaction struct {
	SourceSystem    string    `json:"sourcesystem"`
	SourceTimestamp time.Time `json:"sourcetimestamp"`
	SourceID        string    `json:"sourceid"`
	From            string    `json:"fromaccount"`
	To              string    `json:"toacount"`
	Amount          float64   `json:"amount"`
	Note            string    `json:"note"`
}
