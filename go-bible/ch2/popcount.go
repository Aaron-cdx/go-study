/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 15:44
 * @Motto: Keep thinking, keep coding!
 */

package main

import "fmt"

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

//func init() {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//}

// PopCount returns the population count(number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountTwo(x uint64) int {
	cnt := 0
	for x != 0 {
		x = x & (x - 1)
		fmt.Println(x)
		cnt++
	}
	return cnt
}

func main() {
	//fmt.Println(PopCount(10), PopCountTwo(10))
	//fmt.Println(PopCount(7), PopCount(7))
	fmt.Println(PopCount(3), PopCountTwo(3))
}

// % x means insert a space in to some content.
