/**
 * @Author: caoduanxi
 * @Date: 2021/11/20 15:06
 * @Motto: Keep thinking, keep coding!
 */

package main

import "fmt"

func panicAndRecover() (v int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover()")
		}
	}()
	panic(0)
}

func main() {
	fmt.Println(panicAndRecover())
}
