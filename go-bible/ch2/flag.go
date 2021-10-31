/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 15:12
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s"," ", "separator")

func testFlag(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n{
		fmt.Println()
	}
}

func main() {
	testFlag()
}