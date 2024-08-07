package main

import (
	"fmt"
	"strings"
)

func main () {
	// stings.Index()
	// Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1)

	// strings.Replace()
	// Fungsi ini digunakan untuk replace atau mengganti bagian dari string dengan string tertentu. Jumlah substring yang di-replace bisa ditentukan, apakah hanya 1 string pertama, 2 string, atau kesemuanya.
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newtext1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newtext1)

	var newtext2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newtext2)

	var newtext3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newtext3)

	// strings.Repeat
	// Digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).
	var str1 = strings.Repeat("na", 4)
	fmt.Println(str1)

	// strings.toLower()
	// Mengubah huruf-huruf string menjadi huruf kecil.
	var str2 = strings.ToLower("alAy")
	fmt.Println(str2)

	// strings.toUpper()
	// Mengubah huruf-huruf string menjadi huruf besar.
	var str3 = strings.ToUpper("eat!")
	fmt.Println(str3)
}