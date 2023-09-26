package models

import "time"

type Table struct {
	Name   string     `json:"name,omitempty" bson:"name,omitempty"`
	Header []string   `json:"header,omitempty" bson:"header,omitempty"`
	Rows   [][]string `json:"rows,omitempty" bson:"rows,omitempty"`
}

type Quote struct {
	ID          string    `json:"id,omitempty" bson:"id,omitempty"`
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	ExpiryDate  time.Time `json:"expiryDate,omitempty" bson:"expiryDate,omitempty"`
	TotalAmount float64   `json:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
	Files       []string  `json:"files,omitempty" bson:"files,omitempty"`
	Tables      []Table   `json:"tables,omitempty" bson:"tables,omitempty"`
}


// Quote {
// 	name - name of the quote in text
// 	expiryDate - validity of the quote in Date
// 	status - "valid" or "expired" depending on the quote validity date
// 	totalAmount - total value of quote in number
// 	files - array of files uploaded by the user
// 	tables - array of tables made by the user
// }