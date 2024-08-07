package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fungsi strconv.Atoi()
	// Fungsi ini digunakan untuk konversi data dari tipe string ke int. strconv.Atoi() menghasilkan 2 buah nilai kembalian, yaitu hasil konversi dan error (jika konversi sukses, maka error berisi nil).
	var num = 124
	str, _ := strconv.Itoa(num)

	fmt.Println(str) // "124"
}