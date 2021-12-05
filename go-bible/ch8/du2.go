/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 10:47
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"flag"
	"time"
)

// return the value address
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// ...start background goroutine...
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// print the results periodically
	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}
