package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Latihan 1
	word1 := "Bootcamp"
	word2 := "Digital"
	word3 := "Skill"
	word4 := "Sanbercode"
	word5 := "Golang"

	result := word1 + " " + word2 + " " + word3 + " " + word4 + " " + word5

	fmt.Println(result)


	// Latihan 2
	halo := "Halo Dunia"
	find := "Dunia"
	replace := "Golang"
	
	newHalo := strings.Replace(halo, find, replace, 2 )
	fmt.Println(newHalo)

	// Latihan 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

	var kataKeduaBaru = strings.Replace(kataKedua, "s", "S", 1)
	var kataKetigaBaru = strings.Replace(kataKetiga, "r", "R", 1)
	var kataKeempatBaru = strings.ToUpper(kataKeempat)

	var hasil = kataPertama + " " + kataKeduaBaru + " " + kataKetigaBaru + " " + kataKeempatBaru

	fmt.Println(hasil)

	// Latihan 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";

	angkaPertamaInteger, _ := strconv.Atoi(angkaPertama)
	angkaKeduaInteger, _ := strconv.Atoi(angkaKedua)
	angkaKetigaInteger, _ := strconv.Atoi(angkaKetiga)
	angkaKeempatInteger, _ := strconv.Atoi(angkaKeempat)

	total := angkaPertamaInteger + angkaKeduaInteger + angkaKetigaInteger + angkaKeempatInteger
	fmt.Println("Total: ", total)

	kalimat := "halo halo bandung"
	angka := 2021

	kalimatAwal := "halo halo"
	kalimatBaru := "Hi Hi"

	hasilKalimat := strings.Replace(kalimat, kalimatAwal, kalimatBaru, 2)

	hasilSemua := `"` + hasilKalimat + `"` + " - " + strconv.Itoa(angka)

	fmt.Println(hasilSemua)
}

