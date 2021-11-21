/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 11:14
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"bytes"
	"io"
)

const debug = true

/**
一个不包含任何值的nil接口至和一个刚好包含nil指针的接口值是不同的
*/
func main() {
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// use buf...
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
