/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 21:43
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"time"
)

/**
计算斐波那契数列的第45个元素值
*/
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	// 会打印出动画的效果
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // \r表示不发生回车的动作
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
