/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 23:09
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bitNum = flag.Int("b", 256, "input 256 or 384 or 512 to get sha sum result.\n")

func main() {
	//var arr [3]int // array of 3 integers
	//fmt.Println(arr[0])
	//fmt.Println(arr[len(arr)-1])
	//// self define len of array
	//s := [...]string{"a", "b", "c"}
	//fmt.Println(s)
	flag.Parse()
	calShaSum()
}

func calShaSum() {
	if *bitNum == 256 {
		for _, s := range os.Args[1:] {
			fmt.Printf("%s's sha%d sum result = %x\n", s, *bitNum, sha256.Sum256([]byte(s)))
		}
	} else if *bitNum == 384 {
		for _, s := range os.Args[1:] {
			fmt.Printf("%s's sha%d sum result = %x\n", s, *bitNum, sha512.Sum384([]byte(s)))
		}
	} else {
		for _, s := range os.Args[1:] {
			fmt.Printf("%s's sha%d sum result = %x\n", s, *bitNum, sha512.Sum512([]byte(s)))
		}
	}
	//c1 := sha256.Sum256([]byte("x"))
	//c2 := sha256.Sum256([]byte("X"))
	//fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
