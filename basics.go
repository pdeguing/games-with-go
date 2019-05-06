package main

import "fmt"

func main() {

	/*
		var a int8 // -128, 127
		var b uint* // 0 255
		var c int16
		var d uint16
		var e int // 32 or 64 depending on machine

		// 3.14159
		var f1 float32
		var f2 float64
	*/

	var a uint8 := 250
	var b uint8 := 10
	var result uint8

	result = a + b

	fmt.Println(result)
}
