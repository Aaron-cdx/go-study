/**
 * @Author: caoduanxi
 * @Date: 2021/11/22 23:27
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// NOTE: incorrect!
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

// makeThumbnails3 makes thumbnails of the specified files in parallel
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // ignoring errors
			ch <- struct{}{}
		}(f)
	}
	// wait for goroutines to complete
	for range filenames {
		<-ch // waiting for all goroutine complete work
	}
}

//
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		// 当遇到第一个非nil的error时会直接将error返回到调用方
		// 使得么有一个goroutine去排空errors channel
		// 剩下的goroutine在向这个channel中发送值的时候，就会永远的阻塞下去
		// 并且永远都不会推出，这种情况叫做goroutine泄露
		if err := <-errors; err != nil {
			return err // incorrect: goroutine leak!
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames)) // 带有buffer的chan
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutine
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
