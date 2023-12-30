// Soal 1
package main

import "fmt"

func main() {
	kalimat := "Golang Backend Development"
	tahun := 2021

	// Memanggil fungsi dengan defer
	defer printKalimatDanTahun(kalimat, tahun)

	// Tempatkan kode lain di sini jika diperlukan
}

func printKalimatDanTahun(kalimat string, tahun int) {
	fmt.Printf("Kalimat: %s\n", kalimat)
	fmt.Printf("Tahun: %d\n", tahun)
}


// Soal 2
package main

import (
	"errors"
	"fmt"
)

func kelilingSegitigaSamaSisi(sisi int, tampilkan bool) (string, error) {
	if sisi > 0 {
		hasil := 3 * sisi
		if tampilkan {
			return fmt.Sprintf("Keliling segitiga sama sisi dengan sisi %d cm adalah %d cm", sisi, hasil), nil
		}
		return fmt.Sprint(hasil), nil
	} else {
		if tampilkan {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
		// Menampilkan error beserta panic yang sudah di recover
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic:", r)
			}
		}()
		panic(errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi"))
	}
}

func main() {
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}


// Soal 3
package main

import "fmt"

func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Println("Total Angka:", *total)
}

func main() {
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}

// Soal 4
package main

import "fmt"

var phones = []string{}

func tambahDataPhone(brand string, data *[]string) {
	*data = append(*data, brand)
}

func main() {
	tambahDataPhone("Xiaomi", &phones)
	tambahDataPhone("Asus", &phones)
	tambahDataPhone("IPhone", &phones)
	tambahDataPhone("Samsung", &phones)
	tambahDataPhone("Oppo", &phones)
	tambahDataPhone("Realme", &phones)
	tambahDataPhone("Vivo", &phones)

	fmt.Println("Data Phones:", phones)
}

// Soal 5
package main

import (
	"fmt"
	"math"
)

func luasLingkaran(jariJari float64) int {
	return int(math.Round(math.Pi * jariJari * jariJari))
}

func kelilingLingkaran(jariJari float64) int {
	return int(math.Round(2 * math.Pi * jariJari))
}

func main() {
	jariJari1, jariJari2, jariJari3 := 7.0, 10.0, 15.0

	fmt.Printf("Luas Lingkaran (jari-jari %.2f): %d\n", jariJari1, luasLingkaran(jariJari1))
	fmt.Printf("Keliling Lingkaran (jari-jari %.2f): %d\n", jariJari1, kelilingLingkaran(jariJari1))

	fmt.Printf("Luas Lingkaran (jari-jari %.2f): %d\n", jariJari2, luasLingkaran(jariJari2))
	fmt.Printf("Keliling Lingkaran (jari-jari %.2f): %d\n", jariJari2, kelilingLingkaran(jariJari2))

	fmt.Printf("Luas Lingkaran (jari-jari %.2f): %d\n", jariJari3, luasLingkaran(jariJari3))
	fmt.Printf("Keliling Lingkaran (jari-jari %.2f): %d\n", jariJari3, kelilingLingkaran(jariJari3))
}


// Soal 6
package main

import (
	"flag"
	"fmt"
)

func luasDanKelilingPersegiPanjang(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

func main() {
	// Mendefinisikan flag untuk panjang dan lebar persegi panjang
	panjangPtr := flag.Int("panjang", 0, "Panjang persegi panjang")
	lebarPtr := flag.Int("lebar", 0, "Lebar persegi panjang")
	flag.Parse()

	// Menghitung luas dan keliling persegi panjang
	luas, keliling := luasDanKelilingPersegiPanjang(*panjangPtr
