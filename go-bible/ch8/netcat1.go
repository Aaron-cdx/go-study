/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 21:53
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

/**
conn是两边的共同连接
实际上netcat这边将os.stdin写入，实际上在另一边，那边会读取这边的响应信息
然后进行操作再写回来
这边接收到之后，通过conn和os.Stdout的情况，在重新打印在控制台
*/
//func main() {
//	conn, err := net.Dial("tcp", "localhost:8000")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer conn.Close()
//	//go sendContent(conn)
//	go mustCopy(os.Stdout, conn)
//	mustCopy(conn, os.Stdin)
//}

func sendContent(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.Write(scanner.Bytes())
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
