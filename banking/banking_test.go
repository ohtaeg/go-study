package banking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBankAccount(t *testing.T) {
	// given
	expect := NewBankAccount("ohtaeg")

	// when
	actual := NewBankAccount("ohtaeg")

	// then
	assert.Equal(t, expect, actual)
}

func TestBankAccount_Deposit(t *testing.T) {
	// given
	account := NewBankAccount("ohtaeg")
	expect := 1000

	// when
	account.Deposit(1000)
	actual := account.balance

	// then
	assert.Equal(t, expect, actual)
}

func TestBankAccount_Withdraw(t *testing.T) {
	// given
	account := NewBankAccount("ohtaeg")
	account.Deposit(1000)
	expect := 0

	// when
	err := account.Withdraw(1000)
	actual := account.balance

	// then
	assert.Nil(t, err)
	assert.Equal(t, expect, actual)
}

func TestBankAccount_Withdraw_Error(t *testing.T) {
	// given
	account := NewBankAccount("ohtaeg")
	account.Deposit(1000)

	// when
	err := account.Withdraw(2000)

	// then
	assert.Error(t, err)
}

func TestBankAccount_ChangeOwner(t *testing.T) {
	// given
	account := NewBankAccount("ohtaeg")
	expect := "ohtae"

	// when
	account.ChangeOwner("ohtae")
	actual := account.owner

	// then
	assert.Equal(t, expect, actual)
}
