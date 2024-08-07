package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.ParseFloat()
	// Digunakan untuk konversi string ke numerik desimal dengan lebar data bisa ditentukan.
	var str = "24.12"
	num, _ := strconv.ParseFloat(str, 32)

	fmt.Println(num) // 24.1200008392334
}