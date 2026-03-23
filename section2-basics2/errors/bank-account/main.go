// 🤔 Bank Account
//
// Implement the logic for a bank account. An account has a balance (the amount
// of money in the account) and an overdraft (how far the balance can go below
// zero). Transactions (deposits and withdrawals) are applied to the account
// sequentially. Transactions change the balance; the overdraft stays the same.
//
// You cannot withdraw more than (balance + overdraft). Depositing or
// withdrawing zero is also not allowed.
//
// Example:
//
//   initial state: balance = 100, overdraft = 10
//   transactions: 10, -50, 20
//   result: balance = 80, overdraft = 10
//
// Or:
//
//   initial state: balance = 100, overdraft = 10
//   transactions: -100, -10, -10
//   result: error errInsufficientFunds (insufficient funds)
//
// It is guaranteed that the balance, overdraft, and transaction amounts will
// not overflow int, individually or combined. The overdraft is guaranteed to
// be >= 0.
//
// Sample Input:
//
//   {100 10} [10 -50 20]
//
// Sample Output:
//
//   {80 10}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// errInsufficientFunds signals that the account
// does not have enough funds
// to complete the withdrawal
var errInsufficientFunds = errors.New("insufficient funds")

// errInvalidAmount signals that
// the transaction amount is invalid
var errInvalidAmount = errors.New("invalid transaction amount")

// account represents a bank account
type account struct {
	balance   int
	overdraft int
}

// deposit adds money to the account.
func (acc *account) deposit(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	acc.balance += amount
	return nil
}

// withdraw deducts money from the account.
func (acc *account) withdraw(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}

	maxWithdraw := acc.balance + acc.overdraft
	if amount > maxWithdraw {
		return errInsufficientFunds
	}

	acc.balance -= amount
	return nil
}

// ┌──────────────────────────────────────────┐
// │ do not modify the code below this line   │
// └──────────────────────────────────────────┘

type test struct {
	acc   account
	trans []int
}

var tests = map[string]test{
	"{100 10} [10 -50 20]":  {account{100, 10}, []int{10, -50, 20}},
	"{30 0} [-20 -10]":      {account{30, 0}, []int{-20, -10}},
	"{30 0} [-20 -10 -10]":  {account{30, 0}, []int{-20, -10, -10}},
	"{30 0} [-100]":         {account{30, 0}, []int{-100}},
	"{0 0} [10 20 30]":      {account{0, 0}, []int{10, 20, 30}},
	"{0 0} [10 -10 20 -20]": {account{0, 0}, []int{10, -10, 20, -20}},
	"{20 10} [-20 -10]":     {account{20, 10}, []int{-20, -10}},
	"{20 10} [-20 -10 -10]": {account{20, 10}, []int{-20, -10, -10}},
	"{0 100} [-20 -10]":     {account{0, 100}, []int{-20, -10}},
	"{0 30} [-20 -10]":      {account{0, 30}, []int{-20, -10}},
	"{0 30} [-20 -10 -10]":  {account{0, 30}, []int{-20, -10, -10}},
	"{70 30} [-100 100]":    {account{70, 30}, []int{-100, 100}},
	"{100 10} [10 0 20]":    {account{100, 10}, []int{10, 0, 20}},
}

func main() {
	var err error
	name, err := readString()
	if err != nil {
		log.Fatal(err)
	}
	testCase, ok := tests[name]
	if !ok {
		log.Fatalf("Test case '%v' not found", name)
	}
	for _, t := range testCase.trans {
		if t >= 0 {
			err = testCase.acc.deposit(t)
		} else {
			err = testCase.acc.withdraw(-t)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if err == nil {
		fmt.Println(testCase.acc)
	}
}

// readString reads and returns a string from os.Stdin
func readString() (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
