package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.ParseInt()
	// Digunakan untuk konversi string berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan.
	var str = "124"
	num, _ := strconv.ParseInt(str, 10, 64)

	fmt.Println(num) // 124
}