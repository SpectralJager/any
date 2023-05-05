package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	From   int `json:"from"`
	To     int `json:"to"`
	Amount int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID         int       `json:"id" db:"id"`
	FirstName  string    `json:"firstName" db:"firstName"`
	SecondName string    `json:"secondName" db:"secondName"`
	Number     int64     `json:"number" db:"number"`
	Balance    int64     `json:"balance" db:"balance"`
	CreatedAt  time.Time `json:"createdAt" db:"createdAt"`
}

func NewAccount(firstName string, secondName string) *Account {
	return &Account{
		FirstName:  firstName,
		SecondName: secondName,
		Number:     int64(rand.Intn(111111111)),
		CreatedAt:  time.Now().UTC(),
	}
}
