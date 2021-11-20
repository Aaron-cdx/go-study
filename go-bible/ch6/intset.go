/**
 * @Author: caoduanxi
 * @Date: 2021/11/20 15:16
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// It's zeros value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		// 这里是利用位数，看看位数是否大于，构建bit位置
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// String return the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j) // 这里是按照64进制的数字来计算的
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len return the s length
func (s *IntSet) Len() int {
	return len(s.words)
}

// Remove return remove the value of x
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	// 否则才删除
	word, bit := x/64, uint(x%64)
	// 要把那一位置为0即可,加上那一位是|=(1<<bit)，如果要置为0，取值异或可以把1置为0
	s.words[word] ^= 1 << bit
}

// Clear return remove all element of s
func (s *IntSet) Clear() {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				s.words[i] ^= 1 << uint(j)
			}
		}
	}
}

// Copy return copy of s
func (s *IntSet) Copy() *IntSet {
	var result IntSet
	for _, word := range s.words {
		result.words = append(result.words, word)
	}
	return &result
}

func main() {
	s := &IntSet{}
	s.Add(10)
	s.Add(1)
	s.Add(19)
	s.Add(80)
	fmt.Println(s)
	s.Remove(80)
	s.Remove(19)
	fmt.Println(s)

}
