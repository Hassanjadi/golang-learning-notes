// OPERATOR ARITMATIKA
// adalah operator yang digunakan untuk operasi yang sifatnya perhitungan

// List:
// (+) penjumlahan
// (-) pengurangan
// (*) perkalian
// (/) pembagian
// (%) modulus/sisa hasil pembagian

package main

import (
	"fmt"
)

func main() {
	// operator penjumlahan
	jumlah := 8 + 3
	fmt.Println(jumlah)

	// operator pengurangan
	kurang := 8 - 3
	fmt.Println(kurang)

	// operator perkalian
	kali := 8 * 3
	fmt.Println(kali)

	// operator pembagian
	bagi := 8 / 4
	fmt.Println(bagi)

	// operator modulus
	modulus := 8 % 3
	fmt.Println(modulus)

	// AUGMENTED ASIGNMENT
	// mat a = a + 10 > aug a+=10
	// mat a = a - 10 > aug a-=10
	// mat a = a * 10 > aug a*=10
	// mat a = a / 10 > aug a/=10
	// mat a = a % 10 > aug a%=10

	var angka = 8
	fmt.Println(angka)
	angka += 10
	fmt.Println(angka)

	var angka2 = 5
	fmt.Println(angka2)
	angka2 += 5
	fmt.Println(angka2)

	// UNARY OPERATOR
	// ope ( ++ ) ket a = a + 1
	// ope ( -- ) ket a = a - 1
	// ope ( - ) ket negative
	// ope ( + ) ket positive
	// ope ( ! ) ket negasi/kebalikan dari tipe data boolean

	angkaa := 8
	fmt.Println(angkaa) // 8
	angkaa++
	fmt.Println(angkaa) // 9


	angkaa2 := 5
	fmt.Println(angkaa2) // 5
	angkaa2--
	fmt.Println(angkaa2) // 4
}