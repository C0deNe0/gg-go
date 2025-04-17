package main

import (
	"math/rand"
	"time"
)
//we created this because it is not a good practice to make request to the Account everytime..coz we are not passing the id the number,balance 
//so that's why we use the createAccountRequest to just the needed for creating an account.
type CreateAccountRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}


//This is like a class of the Account we want to create
type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

//and this is the constructor we had created to call that class and intialized with the provided data.
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}
}
