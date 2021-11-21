/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 23:10
 * @Motto: Keep thinking, keep coding!
 */

package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// 如果我们使用了无缓存的channel，则慢的channel将会因为没有人接收
// 而被永远卡住，这种情况称为goroutines泄露，这将是一个Bug
// 和垃圾变量不同，泄露的goroutine不会被自动回收，
// 因此确保每个不再需要的goroutine能正常退出是重要的
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.gopl.io")
	}()
	go func() {
		responses <- request("europe.gopl.io")
	}()
	go func() {
		responses <- request("americas.gopl.io")
	}()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	return hostname
}
