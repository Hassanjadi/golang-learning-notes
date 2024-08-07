package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.FormatInt()
	// Berguna untuk konversi data numerik int64 ke string dengan basis numerik bisa ditentukan sendiri.
	var num = int64(24)
	str := strconv.FormatInt(num, 8)

	fmt.Println(str) // 30
}