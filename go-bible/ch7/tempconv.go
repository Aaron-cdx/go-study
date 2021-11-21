/**
 * @Author: caoduanxi
 * @Date: 2021/11/21 10:59
 * @Motto: Keep thinking, keep coding!
 */

package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

// *celsiusFlag satisfies the flag.Value interface
type celsiusFlag struct {
	Celsius
}

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// 这里因为嵌套了Celsius所以实现了嵌套类的String()方法，所以就具备了输出℃的功能
//func (f *celsiusFlag) String() string {
//	return fmt.Sprintf("%f°C", f.Celsius)
//}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
