/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 17:18
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

/**
针对context包的学习
*/

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		if atomic.LoadInt32(&num) == int32(total) {
			cancelFunc()
		}
	}
	<-ctx.Done()
	fmt.Println("End")
}
