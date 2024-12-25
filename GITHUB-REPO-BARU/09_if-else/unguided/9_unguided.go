//soalPertama
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var beratGr float64

// 	// Input berat parsel dalam gram
// 	fmt.Print("Masukkan berat parsel (gram): ")
// 	fmt.Scan(&beratGr)

// 	// Hitung total berat dalam kg dan sisa dalam gram
// 	beratKg := int(beratGr) / 1000
// 	sisaGram := int(beratGr) % 1000

// 	// Biaya dasar (Rp. 10.000,- per kg)
// 	biayaDasar := beratKg * 10000
// 	tambahanBiaya := 0

// 	// Hitung tambahan biaya jika total berat <= 10kg
// 	if beratKg <= 10 {
// 		if sisaGram >= 500 {
// 			tambahanBiaya = sisaGram * 5
// 		} else {
// 			tambahanBiaya = sisaGram * 15
// 		}
// 	}

// 	// Hitung total biaya pengiriman
// 	totalBiaya := biayaDasar + tambahanBiaya

// 	// Tampilkan hasil perhitungan
// 	fmt.Printf("Total Berat: %d kg %d gram\n", beratKg, sisaGram)
// 	fmt.Printf("Biaya Dasar: Rp %d\n", biayaDasar)
// 	if beratKg <= 10 {
// 		fmt.Printf("Tambahan Biaya: Rp %d\n", tambahanBiaya)
// 	} else {
// 		fmt.Println("Tambahan Biaya: Gratis")
// 	}
// 	fmt.Printf("Total Biaya Pengiriman: Rp %d\n", totalBiaya)
// }

//soalKedua
/*
Jawaban Soal
a. Output atau keluaran dari hasil percabangan dengan inputan 80.1 adalah A dan program yang di minta adalah benar karena nilai A di dapatkan
   Ketika nilai > 80
b. Kesalahan yang pertama adalah variable pada print nilai yang harusnya menggunakan "nmk" bukan "nam" kesalahan selanjutnya adalah dengan mengasih
   range atau batasan nilai agar memunculkan output yang diinginkan. Selanjutnya jika sudah menggunakan if sangat di sarankan selanjutnya menggunakan
   else if
c. Lihat lah code dibawah ini
*/

// package main

// import "fmt"

// func main() {
// 	var nam float64
// 	var nmk string
// 	fmt.Print("Nilai akhir mata kuliah: ")
// 	fmt.Scan(&nam)
// 	if nam > 80 {
// 		nmk = "A"
// 	} else if nam <= 80 && nam > 72.5 {
// 		nmk = "AB"
// 	} else if nam <= 72.5 && nam > 65 {
// 		nmk = "B"
// 	} else if nam <= 65 && nam > 57.5 {
// 		nmk = "BC"
// 	} else if nam <= 57.5 && nam > 50 {
// 		nmk = "C"
// 	} else if nam <= 50 && nam > 40 {
// 		nmk = "D"
// 	} else {
// 		nmk = "E"
// 	}
// 	fmt.Println("Nilai mata kuliah: ", nmk)
// }

//soalKetiga
// package main

// import "fmt"

// func main() {
// 	var b int
// 	fmt.Print("Bilangan: ")
// 	fmt.Scanln(&b)

// 	if b <= 1 {
// 		fmt.Println("Bilangan harus lebih besar dari 1.")
// 	}

// 	fmt.Print("Faktor: ")
// 	faktor := 0
// 	for i := 1; i <= b; i++ {
// 		if b%i == 0 {
// 			fmt.Print(i, " ")
// 			faktor++
// 		}
// 	}
// 	fmt.Println()

// 	prima := b % 2

// 	if prima == 0 {
// 		fmt.Print("Prima : False")
// 	} else {
// 		fmt.Print("Prima : True")
// 	}
// }
