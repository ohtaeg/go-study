package banking

import (
	"errors"
	"fmt"
)

type bankAccount struct {
	owner   string
	balance int
}

// NewBankAccount : function
func NewBankAccount(owner string) *bankAccount {
	bankAccount := bankAccount{owner: owner, balance: 0}
	return &bankAccount
}

// Deposit : Method
func (b *bankAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *bankAccount) Withdraw(amount int) error {
	if b.balance < amount {
		return errors.New("Can't withdraw, you are poor")
	}
	b.balance -= amount
	return nil
}

func (b *bankAccount) ChangeOwner(new string) {
	b.owner = new
}

func (b bankAccount) Balance() int {
	return b.balance
}

func (b bankAccount) Owner() string {
	return b.owner
}

func (b bankAccount) String() string {
	return fmt.Sprint(b.owner, "'s account. Has ", b.balance)
}
