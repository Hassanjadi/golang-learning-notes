// Keyword tipe data bisa digunakan untuk casting, atau konversi antar tipe data. Cara penggunaannya adalah dengan menuliskan tipe data tujuan casting sebagai fungsi, lalu menyisipkan data yang akan dikonversi sebagai parameter fungsi tersebut.

package main

import (
	"fmt"
)

func main () {
	var a float64 = float64(24)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	var str = "Halo"
	var c string = string(str[0])
	fmt.Println(c)
}