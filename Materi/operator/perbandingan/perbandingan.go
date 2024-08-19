package main

import "fmt"

func main () {
	// OPERATOR PERBANDINGAN
	// operator perbandingan digunakan untuk menentukan kebenaran suatu kondisi. Hasilnya berupa nilai boolean, true atau false.

	// > : lebih dari
	// < : kurang dari
	// >= : lebih dari sama dengan
	// <= : kurang dari sama dengan
	// == : sama dengan
	// != : tidak sama dengan

	var angka = 8
	fmt.Println(angka > 5) // true
	fmt.Println(angka < 5) // false
	fmt.Println(angka >= 5) // true
	fmt.Println(angka <= 5) // false
	fmt.Println(angka == 5) // false
	fmt.Println(angka != 5) // true
}