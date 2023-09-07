package nwelastic

import (
	"time"
)

// News describes a document that can be inserted to rethinkdb. Each field is commented with the sources it applies to.
type News struct {
	Id              int64     `json:"id,omitempty"`
	Headline        string    `json:"headline"`
	Body            string    `json:"body,omitempty"`
	Tickers         []string  `json:"tickers,omitempty"`
	Source          string    `json:"source"`
	PublicationTime time.Time `json:"publicationTime"`
	ReceivedTime    time.Time `json:"receivedTime"`
	CreationTime    time.Time `json:"creationTime"` // Override by insert function

	// SEC
	Ciks []int `json:"ciks,omitempty"`
	// SEC
	Link string `json:"link,omitempty"`
}
