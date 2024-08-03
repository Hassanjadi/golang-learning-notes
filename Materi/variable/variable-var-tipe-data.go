// deklarasi dengan var dan juga tipe data

package main

import (
	"fmt"
)

func main() {
	// contoh 1
	var text1 string
	text1 = "ini teks 1"
	fmt.Println(text1)

	// contoh 2
	var text2 string = "ini teks 2"
	text2 = "ini teks 2 diubah"
	fmt.Println(text2)
}