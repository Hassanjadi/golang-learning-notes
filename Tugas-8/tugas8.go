// Soal 1
package main

import (
	"fmt"
	"math"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

func main() {
	segitiga := segitigaSamaSisi{alas: 5, tinggi: 4}
	persegi := persegiPanjang{panjang: 6, lebar: 3}
	silinder := tabung{jariJari: 2.5, tinggi: 8.0}
	kotak := balok{panjang: 4, lebar: 4, tinggi: 4}

	fmt.Println("Perhitungan Bangun Datar:")
	fmt.Printf("Luas Segitiga: %d\n", segitiga.luas())
	fmt.Printf("Keliling Segitiga: %d\n", segitiga.keliling())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegi.luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", persegi.keliling())

	fmt.Println("\nPerhitungan Bangun Ruang:")
	fmt.Printf("Volume Silinder: %.2f\n", silinder.volume())
	fmt.Printf("Luas Permukaan Silinder: %.2f\n", silinder.luasPermukaan())
	fmt.Printf("Volume Kotak: %.2f\n", kotak.volume())
	fmt.Printf("Luas Permukaan Kotak: %.2f\n", kotak.luasPermukaan())
}


// Soal 2
package main

import "fmt"

type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Interface untuk menampilkan data
type displayData interface {
	display() string
}

// Implementasi method display untuk phone
func (p phone) display() string {
	return fmt.Sprintf("Phone: %s %s (%d)\nColors: %v", p.brand, p.name, p.year, p.colors)
}

func main() {
	iphone := phone{
		name:   "iPhone 12",
		brand:  "Apple",
		year:   2021,
		colors: []string{"Silver", "Gold", "Graphite"},
	}

	samsung := phone{
		name:   "Galaxy S21",
		brand:  "Samsung",
		year:   2021,
		colors: []string{"Phantom Gray", "Phantom White", "Phantom Violet"},
	}

	// Menggunakan interface displayData
	fmt.Println("Informasi iPhone:")
	fmt.Println(displayInfo(iphone))

	fmt.Println("\nInformasi Samsung:")
	fmt.Println(displayInfo(samsung))
}

// Fungsi umum untuk menampilkan data dari objek yang mengimplementasikan displayData
func displayInfo(d displayData) string {
	return d.display()
}


// Soal 3
package main

import (
	"fmt"
)

func luasPersegi(sisi int, tampilkan bool) interface{} {
	if sisi > 0 {
		hasil := sisi * sisi
		if tampilkan {
			return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, hasil)
		}
		return hasil
	} else {
		if tampilkan {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}
}

func main() {
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
}


// Soal 4
package main

import "fmt"

func main() {
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	hasil := calculateSum(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua)
	fmt.Println(hasil)
}

func calculateSum(prefix, angkaPertama, angkaKedua interface{}) string {
	// Type assertion untuk mengakses nilai di dalam variabel interface{}
	prefixStr, ok1 := prefix.(string)
	angkaPertamaArr, ok2 := angkaPertama.([]int)
	angkaKeduaArr, ok3 := angkaKedua.([]int)

	if !ok1 || !ok2 || !ok3 {
		return "Error: Tipe data tidak sesuai"
	}

	// Menghitung jumlah dari dua kumpulan angka
	total := 0
	for _, angka := range angkaPertamaArr {
		total += angka
	}
	for _, angka := range angkaKeduaArr {
		total += angka
	}

	// Menghasilkan output yang diinginkan
	return fmt.Sprintf("%s%d + %d + %d + %d = %d", prefixStr, angkaPertamaArr[0], angkaPertamaArr[1], angkaKeduaArr[0], angkaKeduaArr[1], total)
}
