/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 10:46
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"io"
	"time"
)

/**
知识积累
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。
*/

func interfaceExtend() {
	//var w io.Writer
	//w = os.Stdout
	//w = new(bytes.Buffer)
	//w = time.Second // compile error, time.Duration lacks Write method
}

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g "mp3" "WAV"
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
	Resolution() (x, y int)
}
