/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 16:53
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(comma("1234567891"))
	strToIntAndIntToStr()
}

func comma(s string) string {
	buffer := bytes.Buffer{}
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := 0; i < n; i += 3 {
		if i+3 >= n { // if i+3 >= len, means we need append all content
			buffer.WriteString(s[i:])
		} else { // else we need append the three word and comma
			buffer.WriteString(s[i:i+3] + ",")
		}

	}
	return buffer.String()
}

func strToIntAndIntToStr() {
	fmt.Println(strconv.Atoi("100"))                        // string to int
	fmt.Println(strconv.Itoa(100))                          // int to string
	fmt.Println(strconv.ParseInt("10020303003003", 10, 64)) // parse the string to 64bit int
	fmt.Println(strconv.FormatInt(10, 2))
}
