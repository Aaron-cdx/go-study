/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 15:39
 * @Motto: Keep thinking, keep coding!
 */

package bank

import "sync"

var (
	mu       sync.Mutex // guards balance
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock()
	balance3 += amount
	mu.Unlock() // can use defer mu.Unlock() to process unlock
}

func Balance3() int {
	mu.Lock()
	b := balance3
	mu.Unlock()
	return b
}

// 在lock内部执行的都是并发安全的
func Withdraw3(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance3 < 0 {
		return false // insufficient funds
	}
	return true
}

func deposit(amount int) {
	balance3 += amount
}
