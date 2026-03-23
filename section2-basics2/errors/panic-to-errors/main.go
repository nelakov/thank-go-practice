// 🤔 From Panics to Errors
//
// I wrote a program that reads an input string of the form:
//
//   balance/overdraft t1 t2 t3 ... tn
//
// and converts it into an account of type account and a list of transactions
// of type []int.
//
// Unfortunately, I was lazy and did not handle errors explicitly, relying
// instead on catching panics via defer and recover(). Fix the program by
// replacing panics with explicit error returns and handling.
//
// The resulting program must not contain defer or recover(). If the input
// string has an error, the program should print only the error. Otherwise,
// it should print the account and transactions.
//
// Example:
//
//   input:  80/10 10 -20 30
//   output: -> {80 10} [10 -20 30]
//
// Or:
//
//   input:  80/10 10 z 30
//   output: -> strconv.Atoi: parsing "z": invalid syntax
//
// If the input string contains multiple errors, the program should print
// only the first one.
//
// Do not change the error messages — use the same strings that were passed
// to panic() in the original program.
//
// It is guaranteed that the overall format of the input string is valid
// (values are separated by spaces, overdraft is separated from balance by
// a slash, etc.), but individual values may contain errors (as in the
// example above).
//
// The program must never terminate with a panic or os.Exit(1) under any
// circumstances.
//
// Sample Input:
//
//   80/10 10 -20 30
//
// Sample Output:
//
//   -> {80 10} [10 -20 30]

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// do not remove, needed for verification
var _ = errors.As
var _ = reflect.Append
var _ = runtime.Gosched

// account represents a bank account
type account struct {
	balance   int
	overdraft int
}

func main() {
	var (
		acc   account
		trans []int
		err   error
	)

	fmt.Print("-> ")
	acc, trans, err = parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(acc, trans)
}

// parseInput reads an account and a list of transactions from os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)
	if err != nil {
		return account{}, nil, err
	}

	trans, err := parseTransactions(transSrc)
	if err != nil {
		return acc, nil, err
	}

	return acc, trans, nil
}

// readInput returns a string describing the account
// and a slice of strings describing the list of transactions.
// you do not need to modify this function
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount parses an account from a string
// in the format balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return account{}, err
	}

	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return account{}, err
	}

	if overdraft < 0 {
		return account{}, errors.New("expect overdraft >= 0")
	}

	if balance < -overdraft {
		return account{}, errors.New("balance cannot exceed overdraft")
	}

	return account{balance, overdraft}, nil
}

// parseTransactions parses a list of transactions from a string
// in the format [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		trans[idx] = t
	}
	return trans, nil
}
