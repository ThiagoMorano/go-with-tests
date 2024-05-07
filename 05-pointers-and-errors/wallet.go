package pointersanderrors

import (
	"errors"
	"fmt"
)

type Zeny int

func (gp Zeny) String() string {
	return fmt.Sprintf("%d GP", gp)
}

type Wallet struct {
	balance Zeny
}

func (w *Wallet) Balance() Zeny {
	return w.balance
}

func (w *Wallet) Deposit(amount Zeny) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Zeny) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
