/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 11:16
 * @Motto: Keep thinking, keep coding!
 */

package main

import "os"

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// 可以将取消的延时从几百毫秒降低到几十毫秒
func improveDirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token

	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	// read directory
	return nil
}
