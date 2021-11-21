/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 17:44
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"context"
	"fmt"
	"time"
)


// reference: https://www.bilibili.com/video/BV17K411H7iw?from=search&seid=1315949319361049497&spm_id_from=333.337.0.0
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "node01")
	go worker(ctx, "node02")
	go worker(ctx, "node03")
	go worker(ctx, "node04")

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "got the stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
