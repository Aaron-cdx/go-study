/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 15:17
 * @Motto: Keep thinking, keep coding!
 */

package ch9

import "testing"

// 典型的主线程先结束，但是goroutine还没有进行对应操作，导致报错
func TestConcurrency(t *testing.T) {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 10000) }()
	x[9999] = 1
}
