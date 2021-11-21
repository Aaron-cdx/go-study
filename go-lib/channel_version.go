/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 21:37
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	stop <- true
	time.Sleep(5 * time.Second)
}
