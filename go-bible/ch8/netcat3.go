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
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()                    // 这是关闭了读写
	conn.(*net.TCPConn).CloseRead() // 这仅仅是关闭了内部的读，写还是可以继续的
	<-done                          // wait for background goroutine to finish
}
