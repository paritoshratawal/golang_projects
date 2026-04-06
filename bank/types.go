package main

import (
"math/rand"
)

type Account struct{
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Number int64 `json:"number"`
	Balance float64 `json:"balance"`
}

func NewAccount(FirstName, LastName string) *Account{
	return &Account{
		ID: int(rand.Intn(1000)),
		FirstName: FirstName,
		LastName: LastName,
		Number: int64(rand.Intn(1000000000)),
		Balance: 0.0,
	}
}