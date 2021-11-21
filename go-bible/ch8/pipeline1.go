/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 23:03
 * @Motto: Keep thinking, keep coding!
 */

package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer in main goroutine
	for {
		fmt.Println(<-squares)
	}
}
