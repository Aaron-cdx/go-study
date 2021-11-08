/**
 * @Author: caoduanxi
 * @Date: 2021/11/7 23:12
 * @Motto: Keep thinking, keep coding!
 */
/**
判断slice是否为空，使用len(s) == 0会比判
*/
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	//s := []int{0, 1, 2, 3, 4, 5}
	//reverse(s[:])
	//fmt.Println(s)

	//fmt.Println(uint(-20))
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6}))
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[i:]
}

func reverseII(s *[]int) {
	// 使用数组指针代替slice
	i, j := 0, len(*s)-1
	for i < j {
		// 这个就是数组指针，当我们需要使用的时候，可以使用*s将指针数组转化为数组
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1
	}
}

func rotate(s []int) []int {
	l := len(s)
	res := make([]int, l)
	// 完全倒序
	for i := 0; i < l; i++ {
		res[l-i-1] = s[i]
	}
	return res
}
