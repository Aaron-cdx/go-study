/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 15:27
 * @Motto: Keep thinking, keep coding!
 */
// Package bank provides a concurrency-safe bank with one account
package bank

import "fmt"

var deposits = make(chan int)    // send amount to deposit
var balances = make(chan int)    // receive balance
var withdraw = make(chan string) // receive the balance whether sufficient

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func WithDraw() {
	if Balance() <= 0 {
		withdraw <- "balance not sufficient"
	}
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case s := <-withdraw:
			fmt.Println(s)
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touched this cake again
	}
}
