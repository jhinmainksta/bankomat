package service

import "errors"

type Account struct {
	balance float64
}

type Amount struct {
	Amount float64
}

func (a *Account) Deposit(amount float64) error {
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return errors.New("not enough рублей")
	}
	a.balance -= amount
	return nil
}

func (a Account) GetBalance() float64 {
	return a.balance
}
