/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 10:53
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"flag"
	"fmt"
	"time"
)

/**
flag.Value可以帮助行标记定义新的符号
String方法格式化标记的值用在命令行帮助消息中，这样每一个flag.Value也是一个fmt.Stringer
Set方法就是解析它的字符串参数并且更新为标记变量的值。
 */
var period = flag.Duration("period", 1*time.Second, "sleep period")


func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}



