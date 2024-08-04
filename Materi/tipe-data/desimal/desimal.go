// tipe data Numerik Desimal
// float32 dan float64. Perbedaaan tipe data berada di lebar cakupan nilai desimal yang bisa ditampung
package main

import (
	"fmt"
)

func main()  {
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
}