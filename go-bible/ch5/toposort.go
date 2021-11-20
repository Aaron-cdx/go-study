/**
 * @Author: caoduanxi
 * @Date: 2021/11/20 14:17
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// topoSort的主要的作用就是利用每一门课向上追溯一定会找到第一门课程！所以无论从哪一门课程开始一定会找到第一门课程
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
/**
为什么下面的dir需要重新声明：
	问题的原因在于循环变量的作用域，在上面的程序中，for循环语句引入了新的语法块，循环变量
dir在这个词块中被声明，在该循环中生成的所有函数都共享相同的循环变量，需要注意，函数值中记录的是循环变量的
内存地址，而不是循环变量某一个时刻的值，以dir为例子，后续的迭代会不断的更新dir的值，当删除操作执行完时，for循环已经完成
dir中存储的值等于最后一次迭代的值，这意味着每次对os.RemoveAll()的调用删除都是相同的目录

通常为了解决这个问题，我们会引入一个与循环变量同名的局部变量，作为循环变量的副本
 */
func catchDfsVariable(){
	/*var rmdirs []func()
	for _,d := range tempDirs(){
		dir := d // 这里的变量需要重新声明
		os.MkdirAll(dir,0755) // create parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}*/
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}
