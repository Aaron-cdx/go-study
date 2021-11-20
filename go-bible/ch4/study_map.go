/**
 * @Author: caoduanxi
 * @Date: 2021/11/8 23:37
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
不能够对map中元素进行取址操作，原因是因为随着元素数量的增长map会重新
分配更大的内存空间，从而可能导致之前的地址无效。
*/

// 判断两个map是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return false
}

func dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(-1)
	}
}
