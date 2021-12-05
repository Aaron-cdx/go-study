/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 22:48
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}) // 这里是因为done在这里，所以可以维持通信下去
	go func() {
		io.Copy(os.Stdout, conn) // ignoring errors
		log.Println("done")      // 一旦done被放入元素，则整体结束
		done <- struct{}{}       // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()                    // 这是关闭了读写
	conn.(*net.TCPConn).CloseRead() // 这仅仅是关闭了内部的读，写还是可以继续的
	<-done                          // wait for background goroutine to finish
}
