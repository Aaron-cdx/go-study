/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 21:48
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		handleConn(conn)
	}
}

// fist version of handleConn
/*func handleConn(c net.Conn) {
	defer c.Close()
	//data := make([]byte, 1024)
	//_, err := c.Read(data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(data)
	for {
		// 辅助记忆 2006年01月02日 下午15点04分05秒go语言诞生
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}*/

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\t%s", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// ignoring potential errors from input.Err()
	c.Close()
}
