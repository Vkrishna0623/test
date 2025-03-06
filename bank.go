package test

import "errors"

type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
	if amount <= 0 {
		return b.Balance, errors.New("withdrawal amount must be greater than zero")
	}
	if amount > b.Balance {
		return b.Balance, errors.New("insufficient funds")
	}
	b.Balance -= amount
	return b.Balance, nil
}
