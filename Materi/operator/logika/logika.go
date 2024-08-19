package main

import (
	"fmt"
)

func main() {
	// OPERATOR LOGIKA
	// Operator ini digunakan untuk mencari benar tidaknya kombinasi data bertipe bool (bisa berupa variabel bertipe bool, atau hasil dari operator perbandingan).

	// && : dan
	// || : atau
	// ! : negasi/nilai kebalikan

	var a = true
	var b = false
	var c = true
	var d = false
	
	fmt.Println(a && c) // true
	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
	fmt.Println(b || d) // false
	fmt.Println(!b && !d) // true
	fmt.Println(!a && b) // false
}