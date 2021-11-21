/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 23:22
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	calculateTime(goroutineAdd)
	calculateTime(directAdd)
}

func goroutineAdd(n int) int {
	resultChan := make(chan int, 3)
	var calSum func(n, divider int, c chan int)
	calSum = func(start, end int, c chan int) {
		sum := 0
		for i := start; i <= end; i++ {
			sum += i
		}
		c <- sum
	}
	go calSum(1, n/3, resultChan)
	go calSum(n/3+1, n*2/3, resultChan)
	go calSum(n*2/3+1, n, resultChan)
	return <-resultChan + <-resultChan + <-resultChan
}

func directAdd(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func calculateTime(f func(n int) int) {
	now := time.Now()
	f(100000000000)
	fmt.Printf("func execute spend time:%dms\n", time.Since(now))
}
