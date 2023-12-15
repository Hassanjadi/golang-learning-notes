// Soal 1
package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        if i%2 == 1 { // angka ganjil
            fmt.Printf("%d - Santai\n", i)
        } else if i%2 == 0 { // angka genap
            fmt.Printf("%d - Berkualitas\n", i)
        }
        if i%3 == 0 && i%2 == 1 { // kelipatan 3 dan angka ganjil
            fmt.Printf("%d - I Love Coding\n", i)
        }
    }
}

// Soal 2
package main

import "fmt"

func main() {
    for i := 1; i <= 7; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}

// Soal 3
package main

import "fmt"

func main() {
    var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

    // Menggunakan slicing untuk memilih bagian yang diinginkan
    hasil := kalimat[2:7]

    fmt.Println(hasil)
}


// Soal 4
package main

import "fmt"

func main() {
    var sayuran = []string{}

    // Menambahkan data ke variabel sayuran
    sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

    // Menampilkan data dengan loop dan memberikan angka di depannya
    for i, s := range sayuran {
        fmt.Printf("%d. %s\n", i+1, s)
    }
}

// Soal 5
package main

import "fmt"

func main() {
    var satuan = map[string]int{
        "panjang": 7,
        "lebar":   4,
        "tinggi":  6,
    }

    // Menampilkan nilai dari variabel menggunakan loop
    for k, v := range satuan {
        fmt.Printf("%s = %d\n", k, v)
    }

    // Menghitung volume balok
    volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
    fmt.Printf("Volume balok = %d\n", volume)
}
