/**
 * @Author: caoduanxi
 * @Date: 2021/11/20 12:23
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"strings"
)

/**
函数值也是不可比较的，所以也不可以作为map的key
*/

// 编写函数expand,将s中的"foo"替换为f("foo")的返回值
func expand(s string, f func(string) string) string {
	all := strings.ReplaceAll(s, "foo", f("foo"))
	return all
}

func main() {
	s := expand("food,hello world!", func(s string) string {
		s = "hello cdx"
		return s
	})
	fmt.Println(s)
}
