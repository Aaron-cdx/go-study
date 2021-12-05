/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 15:37
 * @Motto: Keep thinking, keep coding!
 */

package bank

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

func Deposit2(amount int) {
	sema <- struct{}{} // acquire token
	balance += amount
	<-sema // release token
}

func Balance2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}


