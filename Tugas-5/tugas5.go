// Soal 1
package main

import "fmt"

func luasPersegiPanjang(panjang, lebar float64) float64 {
    return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar float64) float64 {
    return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi float64) float64 {
    return panjang * lebar * tinggi
}

func main() {
    panjang := 12.0
    lebar := 4.0
    tinggi := 8.0

    luas := luasPersegiPanjang(panjang, lebar)
    keliling := kelilingPersegiPanjang(panjang, lebar)
    volume := volumeBalok(panjang, lebar, tinggi)

    fmt.Println("Luas Persegi Panjang:", luas)
    fmt.Println("Keliling Persegi Panjang:", keliling)
    fmt.Println("Volume Balok:", volume)
}

// Soal 2
package main

import "fmt"

func introduce(nama, jenisKelamin, pekerjaan, usia string) string {
    var panggilan string

    // Menentukan panggilan berdasarkan jenis kelamin
    if jenisKelamin == "laki-laki" {
        panggilan = "Pak"
    } else if jenisKelamin == "perempuan" {
        panggilan = "Bu"
    } else {
        panggilan = ""
    }

    // Membentuk kalimat introduksi
    introduction := fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggilan, nama, pekerjaan, usia)

    return introduction
}

func main() {
    john := introduce("John", "laki-laki", "penulis", "30")
    fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

    sarah := introduce("Sarah", "perempuan", "model", "28")
    fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
}

// Soal 3
package main

import "fmt"

func buahFavorit(nama string, buah ...string) string {
    // Membentuk kalimat berdasarkan data yang diberikan
    kalimat := fmt.Sprintf("Halo, nama saya %s dan buah favorit saya adalah %q", nama, buah)

    return kalimat
}

func main() {
    buah := []string{"semangka", "jeruk", "melon", "pepaya"}
    buahFavoritJohn := buahFavorit("John", buah...)

    fmt.Println(buahFavoritJohn)
    // Output: Halo, nama saya John dan buah favorit saya adalah ["semangka", "jeruk", "melon", "pepaya"]
}

// Soal 4
package main

import (
	"fmt"
)

var dataFilm = []map[string]string{}

func tambahDataFilm() func(string, string, string, string) {
	return func(judul, durasi, genre, tahun string) {
		film := map[string]string{
			"judul":  judul,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, film)
	}
}

func main() {
	tambahDataFilm := tambahDataFilm()

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}