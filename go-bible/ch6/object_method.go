/**
 * @Author: caoduanxi
 * @Date: 2021/11/20 15:12
 * @Motto: Keep thinking, keep coding!
 */

package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := &Point{1, 2}
	p.ScaleBy(0.9)
	fmt.Println(p)
}
