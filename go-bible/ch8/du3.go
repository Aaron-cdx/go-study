/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 11:02
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"flag"
	"path/filepath"
	"sync"
)

func main() {
	// determine roots
	// ...start background goroutine...
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// traverse each root of the file tree in parallel
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDirII(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// select loop
}

func walkDirII(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDirII(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}
