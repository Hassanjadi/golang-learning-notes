package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.FormatFloat()
	// Berguna untuk konversi data bertipe float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data bisa ditentukan.
	var num = float64(24.12)
	str := strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str) // 24.120000
}