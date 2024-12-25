// MODUL 4 SOAL 1
// PROGRAM UNTUK MENGHITUNG HARGA TOTAL SETELAH DIDISKON BERAPA PERSEN TERTENTU
// package main

// import "fmt"

// func main() {
// 	var hargaAwal, diskon int

// 	// Input harga awal dan diskon yang mau ditetapkan
// 	fmt.Scan(&hargaAwal)
// 	fmt.Scan(&diskon)

// 	// Menghitung jumlah dengan diskon
// 	jumlahDiskon := float64(hargaAwal) * float64(diskon) / 100
// 	total := float64(hargaAwal) - jumlahDiskon

// 	// Output harga akhir setelah dipakaikan diskon
// 	fmt.Printf("%.0f\n", total)
// }

// MODUL 4 SOAL 2
// PROGRAM UNTUK MENGHITUNG BERAT BADAN BERDASARKAN BMI DAN TINGGI BADAN SESEORANG
// package main

// import "fmt"

// func main() {
// 	var bmi, tinggi, beratBadan float64

// 	// Input bmi dan tinggi badan
// 	fmt.Scan(&bmi)
// 	fmt.Scan(&tinggi)

// 	// Menghitung berat badan berdasarkan bmi dan tinggi badan yang sudah di input dan scan
// 	beratBadan = bmi * (tinggi * tinggi)

// 	// mengeluarkan hasil akhir dari pengoprasian beratBadan
// 	fmt.Printf("%.0f\n", beratBadan)
// }

// MODUL 4 SOAL 3
// PROGRAM UNTUK MENGHITUNG SISI TERPANJANG SUATU SEGITIGA DENGAN PYTAGORAS
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x1, y1, x2, y2, x3, y3 float64

// 	//input koordinat titik a, b dan c
// 	fmt.Scan(&x1, &y1)
// 	fmt.Scan(&x2, &y2)
// 	fmt.Scan(&x3, &y3)

// 	//Menghitung jarak antara titik menggunakan formula jarak
// 	ab := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
// 	bc := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
// 	ac := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

// 	//Menemukan sisi paling panjang
// 	longest := ab
// 	if bc > longest {
// 		longest = bc
// 	}
// 	if ac > longest {
// 		longest = ac
// 	}

// 	//Output sisi paling panjang dengan 2 desimal di belakan koma
// 	fmt.Printf("%.2f\n", longest)
// }

