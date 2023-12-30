package main

import (
	"fmt"
)

func main() {
	segitiga := geometri.segitigaSamaSisi{alas: 5, tinggi: 4}
	persegi := geometri.persegiPanjang{panjang: 6, lebar: 3}
	silinder := geometri.tabung{jariJari: 2.5, tinggi: 8.0}
	kotak := geometri.balok{panjang: 4, lebar: 4, tinggi: 4}

	fmt.Println("Perhitungan Bangun Datar:")
	fmt.Printf("Luas Segitiga: %d\n", segitiga.Luas())
	fmt.Printf("Keliling Segitiga: %d\n", segitiga.Keliling())
	fmt.Printf("Luas Persegi Panjang: %d\n", persegi.Luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", persegi.Keliling())

	fmt.Println("\nPerhitungan Bangun Ruang:")
	fmt.Printf("Volume Silinder: %.2f\n", silinder.Volume())
	fmt.Printf("Luas Permukaan Silinder: %.2f\n", silinder.LuasPermukaan())
	fmt.Printf("Volume Kotak: %.2f\n", kotak.Volume())
	fmt.Printf("Luas Permukaan Kotak: %.2f\n", kotak.LuasPermukaan())
}
