package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	fund float64
	mux  sync.Mutex
}

func (a *Account) Balance() float64 {
	a.mux.Lock()
	defer a.mux.Unlock()

	return a.fund
}

func (a *Account) Credit(amount float64) {
	a.mux.Lock()
	defer a.mux.Unlock()

	a.fund += amount
}

func (a *Account) Withdraw(amount float64) (float64, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.fund < amount {
		return -1, errors.New("Insufficient funds")
	}

	a.fund -= amount
	return a.fund, nil
}

func main() {
	account := Account{fund: 500.0}

	account.Credit(100.0)
	fmt.Println(account.Balance())
	balance, err := account.Withdraw(650)
	if err == nil {
		fmt.Println(balance)
	} else {
		fmt.Println(err)
	}
	account.Credit(100.0)
	fmt.Println(account.Balance())
}
