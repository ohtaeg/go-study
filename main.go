package main

import (
	"fmt"
	"github.com/ohtaeg/job-scrapper/banking"
	"log"
)

func main() {
	account := banking.NewBankAccount("ohtae")
	log.Println(account)

	account.Deposit(1000)
	fmt.Println(account.Balance())

	err := account.Withdraw(1001)
	if err != nil {
		// log.Fatalln(err) // kill program
		log.Println(err)
	}
	fmt.Println(account.Owner(), account.Balance())

	account.ChangeOwner("ohtaeg")
	fmt.Println(account.Owner(), account.Balance())
}
