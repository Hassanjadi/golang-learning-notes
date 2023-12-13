// Soal 1
package main

import "fmt"

func main() {
	bootcamp := "Bootcamp"
	digital := "Digital"
	skill := "Skill"
	sanbercode := "Sanbercode"
	golang := "Golang"

	kalimat := fmt.Sprintf("%s %s %s %s %s", bootcamp, digital, skill, sanbercode, golang)
	fmt.Println(kalimat)
}


// Soal 2
package main

import (
	"fmt"
	"strings"
)

func main() {
	halo := "Halo Dunia"

	// Mengganti kata "Dunia" menjadi "Golang"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println(halo)
}

// Soal 3
package main

import "fmt"

func main() {
    kataPertama := "saya"
    kataKedua := "senang"
    kataKetiga := "belajar"
    kataKeempat := "golang"

    // Menggabungkan variabel-variabel dan mencetak dengan format tertentu
    hasilGabungan := fmt.Sprintf("%s %s %s %s", kataPertama, kataKedua, kataKetiga, kataKeempat)
    
    // Mengonversi karakter pertama kataKeempat menjadi huruf besar
    hasilGabungan = hasilGabungan[:len(hasilGabungan)-len(kataKeempat)] + strings.ToUpper(string(kataKeempat[0])) + hasilGabungan[len(hasilGabungan)-len(kataKeempat)+1:]

    fmt.Println(hasilGabungan)
}

// Soal 4
package main

import (
	"fmt"
	"strconv"
)

func main() {
	angkaPertama := "8"
	angkaKedua := "5"
	angkaKetiga := "6"
	angkaKeempat := "7"

	// Konversi string ke integer
	angkaIntPertama, _ := strconv.Atoi(angkaPertama)
	angkaIntKedua, _ := strconv.Atoi(angkaKedua)
	angkaIntKetiga, _ := strconv.Atoi(angkaKetiga)
	angkaIntKeempat, _ := strconv.Atoi(angkaKeempat)

	// Jumlahkan semua angka
	total := angkaIntPertama + angkaIntKedua + angkaIntKetiga + angkaIntKeempat

	fmt.Println("Jumlah: ", total)
}


// Soal 5
package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "halo halo bandung"
	angka := 2021

	// Memanipulasi kalimat
	potonganKalimat := strings.Split(kalimat, " ")
	hasilKalimat := "Hi Hi " + potonganKalimat[2]

	// Menggabungkan hasil kalimat dengan angka
	outputAkhir := hasilKalimat + " - " + strconv.Itoa(angka)

	fmt.Println(outputAkhir)
}

