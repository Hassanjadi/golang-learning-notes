package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.FormatBool()
	// Digunakan untuk konversi bool ke string.
	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str) // true
}