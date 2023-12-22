// Soal 1
package main

import "fmt"

type buah struct {
	Nama        string
	Warna       string
	AdaBijinya  bool
	Harga       int
}

func main() {
	nanas := buah{"Nanas", "Kuning", false, 9000}
	jeruk := buah{"Jeruk", "Oranye", true, 8000}
	semangka := buah{"Semangka", "Hijau & Merah", true, 10000}
	pisang := buah{"Pisang", "Kuning", false, 5000}

	fmt.Println("Data Buah:")
	fmt.Println("1. Nama:", nanas.Nama, "| Warna:", nanas.Warna, "| Ada Bijinya:", nanas.AdaBijinya, "| Harga:", nanas.Harga)
	fmt.Println("2. Nama:", jeruk.Nama, "| Warna:", jeruk.Warna, "| Ada Bijinya:", jeruk.AdaBijinya, "| Harga:", jeruk.Harga)
	fmt.Println("3. Nama:", semangka.Nama, "| Warna:", semangka.Warna, "| Ada Bijinya:", semangka.AdaBijinya, "| Harga:", semangka.Harga)
	fmt.Println("4. Nama:", pisang.Nama, "| Warna:", pisang.Warna, "| Ada Bijinya:", pisang.AdaBijinya, "| Harga:", pisang.Harga)
}


// Soal 2
package main

import "fmt"

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func main() {
	triangle := segitiga{alas: 5, tinggi: 3}
	square := persegi{sisi: 4}
	rectangle := persegiPanjang{panjang: 6, lebar: 2}

	fmt.Println("Luas Segitiga:", triangle.luas())
	fmt.Println("Luas Persegi:", square.luas())
	fmt.Println("Luas Persegi Panjang:", rectangle.luas())
}


// Soal 3
package main

import "fmt"

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) tambahWarna(warna string) {
	p.colors = append(p.colors, warna)
}

func main() {
	iphone := phone{name: "iPhone", brand: "Apple", year: 2022, colors: []string{"Silver", "Gold"}}
	samsung := phone{name: "Galaxy", brand: "Samsung", year: 2022, colors: []string{"Black", "Blue"}}

	fmt.Println("Warna iPhone:", iphone.colors)
	fmt.Println("Warna Samsung:", samsung.colors)

	iphone.tambahWarna("Rose Gold")
	samsung.tambahWarna("White")

	fmt.Println("Warna iPhone setelah tambah:", iphone.colors)
	fmt.Println("Warna Samsung setelah tambah:", samsung.colors)
}


// Soal 4
package main

import "fmt"

type movie struct {
	Title    string
	Genre    string
	Duration int
	Year     int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	film := movie{
		Title:    title,
		Genre:    genre,
		Duration: duration,
		Year:     year,
	}

	*dataFilm = append(*dataFilm, film)
}

func main() {
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	fmt.Println("Data Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1, film.Title)
		fmt.Printf("   Genre: %s\n", film.Genre)
		fmt.Printf("   Duration: %d minutes\n", film.Duration)
		fmt.Printf("   Year: %d\n", film.Year)
		fmt.Println()
	}
}
