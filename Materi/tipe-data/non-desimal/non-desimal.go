// tipe data numerik Non-Desimal
// uint: tipe data untuk bilangan cacah (bilangan positif)
// int: tipe data untuk bilangan bulat (bilangan negatif dan positif)
package main

import (
	"fmt"
)

func main()  {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
}