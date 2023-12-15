// Soal 1
package main

import (
	"fmt"
	"strconv"
)

func main() {.
	// Deklarasi variabel
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// Konversi string ke integer
	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	// Perhitungan luas dan keliling persegi panjang
	luasPersegiPanjang = panjang * lebar
	kelilingPersegiPanjang = 2 * (panjang + lebar)

	// Perhitungan luas segitiga
	luasSegitiga = 0.5 * alas * tinggi

	// Menampilkan hasil perhitungan
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)
}

// Soal 2
package main

import "fmt"

func main() {
	// Deklarasi variabel
	var nilaiJohn = 80
	var nilaiDoe = 50

	// Menentukan indeks nilai John
	fmt.Print("Indeks Nilai John: ")
	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// Menentukan indeks nilai Doe
	fmt.Print("Indeks Nilai Doe: ")
	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

// Soal 3
package main

import "fmt"

func main() {
	// Ganti tanggal, bulan, dan tahun sesuai dengan tanggal lahir Anda
	var tanggal = 15
	var bulan = 12
	var tahun = 1990

	// Switch case untuk menentukan nama bulan
	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	// Menampilkan output
	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)
}

// Soal 4
package main

import "fmt"

func main() {
	// Ganti tahun kelahiran sesuai dengan tahun kelahiran Anda
	var tahunKelahiran = 1990

	// Menentukan generasi dengan menggunakan conditional
	var generasi string

	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		generasi = "Baby Boomer"
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		generasi = "Generasi X"
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		generasi = "Generasi Y (Millennials)"
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Generasi tidak terdefinisi"
	}

	// Menampilkan output
	fmt.Printf("Anda termasuk dalam generasi: %s\n", generasi)
}
