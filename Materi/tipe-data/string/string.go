// digunakan untuk mencari posisi index sebuah string
package main

import (
	"fmt"
)

func main() {
	// dengan tanda quote ("")
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	// dengan tanda grave accent/backticks (`)
	// contoh 1
	var messages = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang"`
	fmt.Println(messages)

	// contoh 2
	var name = "John Doe"
	var age = "28"
	var sentence = `halo nama saya "` + name + `" dan berumur "`+ age +`"`

	fmt.Println()
	fmt.Println(sentence)
}