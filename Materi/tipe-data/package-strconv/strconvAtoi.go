package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.Itoa()
	// Merupakan kebalikan dari strconv.Atoi, berguna untuk konversi int ke string.
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) //124
	}
}