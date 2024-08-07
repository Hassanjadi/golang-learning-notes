package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.ParseBool()
	// Digunakan untuk konversi string ke bool.
	var str = "true"
	var bul, _ = strconv.ParseBool(str)

	fmt.Println(bul) // true
}