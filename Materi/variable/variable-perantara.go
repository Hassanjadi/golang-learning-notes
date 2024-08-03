// deklarasi dengan menggunakan perantara ":="

package main

import (
	"fmt"
)

func main() {
	// contoh 1
	text1 := "ini text 1"
	text1 = "ini teks 1 diubah"
	fmt.Println(text1)

	// contoh 2
	text2 := "ini teks 2"
	fmt.Println(text2)
}