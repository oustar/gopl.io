// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type WithDraw struct {
	n  int
	ok chan bool
}

var deposits = make(chan int)       // send amount to deposit
var balances = make(chan int)       // receive balance
var withdraws = make(chan WithDraw) // withdraw amount

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- WithDraw{amount, ch}
	ok := <-ch
	close(ch)
	return ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wdinfo := <-withdraws:
			if wdinfo.n > balance {
				wdinfo.ok <- false
			}
			balance -= wdinfo.n
			wdinfo.ok <- true
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
