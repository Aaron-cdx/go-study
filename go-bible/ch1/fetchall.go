/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 14:12
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/**
这里还涉及如果是goroutine blocked的情况，如果有goroutine卡在了某个网站获取不到数据，此时会进行具体blocked，需要考虑如何退出
 */
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch, if any goroutine not process completed, in here will blocked
	}
	fmt.Printf("%.2fs elaspsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s, %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) // format a string to toString
}
