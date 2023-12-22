// Soal 1
package main

import "fmt"

var luasLingkaran float64
var kelilingLingkaran float64

func hitungLingkaran(jariJari *float64) {
	const pi = 3.14
	luasLingkaran = pi * (*jariJari) * (*jariJari)
	kelilingLingkaran = 2 * pi * (*jariJari)
}

func main() {
	jariJariLingkaran := 5.0
	hitungLingkaran(&jariJariLingkaran)
	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)
}


// Soal 2
package main

import "fmt"

var sentence string

func introduce(name, gender, occupation, age string) {
	var prefix string

	if gender == "laki-laki" {
		prefix = "Pak "
	} else {
		prefix = "Bu "
	}

	sentence = fmt.Sprintf("%s%s adalah seorang %s yang berusia %s tahun", prefix, name, occupation, age)
}

func main() {
	introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}


// Soal 3
package main

import "fmt"

var buah = []string{}

func tambahBuah(namaBuah string) {
	buah = append(buah, namaBuah)
}

func main() {
	tambahBuah("Jeruk")
	tambahBuah("Semangka")
	tambahBuah("Mangga")
	tambahBuah("Strawberry")
	tambahBuah("Durian")
	tambahBuah("Manggis")
	tambahBuah("Alpukat")

	fmt.Println("Daftar Buah:")
	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}
}


// Soal 4
package main

import (
	"fmt"
)

var dataFilm = []map[string]string{}

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := make(map[string]string)
	film["Title"] = title
	film["Duration"] = duration
	film["Genre"] = genre
	film["Year"] = year

	*dataFilm = append(*dataFilm, film)
}

func main() {
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// Menampilkan data
	fmt.Println("Data Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1, film["Title"])
		fmt.Printf("   Duration: %s\n", film["Duration"])
		fmt.Printf("   Genre: %s\n", film["Genre"])
		fmt.Printf("   Year: %s\n", film["Year"])
		fmt.Println()
	}
}
