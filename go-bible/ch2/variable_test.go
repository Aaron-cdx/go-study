/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 15:04
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"testing"
)

func TestPointerAdd(t *testing.T){
	//testFlag()
}

// do the pointer add process
func incr(p *int) int{
	*p++
	return *p
}

