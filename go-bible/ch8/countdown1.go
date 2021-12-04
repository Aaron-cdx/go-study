/**
 * @Author: caoduanxi
 * @Date: 2021/12/4 22:40
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"os"
	"time"
)

// 基于select的多路复用
func main() {
	//fmt.Println("Commencing countdown.")
	//tick := time.Tick(1 * time.Second) // 倒计时时间戳，可以实现倒计时功能，等同于sleep(time)
	//for countdown := 10; countdown > 0; countdown-- {
	//	fmt.Println(countdown)
	//	<-tick
	//}
	//launch()

	//abort()

	selectBuffer()
}

func launch() {
	fmt.Println("Rocket launching~")
}

func abort() {
	abort := make(chan struct{}) // abort, user can abort the launch
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	// multiplex,多路复用机制
	select {
	case <-time.After(10 * time.Second): // 时间结束自己会发送到C中，所以可以接受到这个10秒的信号
		// do nothing
		fmt.Println("time already elapse 10 seconds, rocket launched.")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

// 或交替为空或者为满，只有一个case可以进行下去，无论i是奇数还是偶数，打印的都是0 2 4 6 8
func selectBuffer() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		// 当0的时候只能够放入，所以放入0，当1的时候，取出和放入都可以，所以随机选择了第一个1
		// 然后就输出了0，随后为2的时候，放入2(好奇为什么随机选择不是选择放入呢，因为必须要取出，不能够放入)
		select {
		case x := <-ch:
			fmt.Println(x) // 0 2 4 6 8
		case ch <- i:
		}
	}
}
