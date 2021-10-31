/**
 * @Author: caoduanxi
 * @Date: 2021/10/30 15:36
 * @Motto: Keep thinking, keep coding!
 */

package main

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
