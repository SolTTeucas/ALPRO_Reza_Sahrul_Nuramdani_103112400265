//soalPertama
// package main

// import "fmt"

// func main() {
// 	// Input: Bilangan bulat positif
// 	var number int
// 	fmt.Print("Masukkan bilangan bulat positif: ")
// 	fmt.Scan(&number)

// 	// Cek apakah bilangan positif
// 	if number <= 0 {
// 		fmt.Println("Masukkan bilangan bulat positif.")
// 		return
// 	}

// 	// Menghitung banyaknya digit
// 	digitCount := 0
// 	for number > 0 {
// 		number /= 10 // Menghapus digit terakhir
// 		digitCount++ // Menambah jumlah digit
// 	}

// 	// Output: Banyaknya digit
// 	fmt.Printf("Banyaknya digit adalah %d\n", digitCount)
// }

//soalKedua
// package main

// import "fmt"

// func main() {
// 	// Input: Bilangan desimal
// 	var number float64
// 	fmt.Print("Masukkan bilangan desimal: ")
// 	fmt.Scan(&number)

// 	// Pembulatan ke atas (ceiling)
// 	roundedUp := float64(int(number))
// 	if number > roundedUp {
// 		roundedUp++
// 	}

// 	// Repeat-until style: menambahkan 0.1 setiap iterasi
// 	// Mulai dari nilai input yang sudah dibulatkan
// 	for number < roundedUp {
// 		// Menampilkan nilai setelah penambahan
// 		number += 0.1
// 		fmt.Printf("%.1f\n", number)
// 	}
// }

//soalKetiga
// package main

// import "fmt"

// func main() {
// 	// Input: Target donasi
// 	var target int
// 	fmt.Print("Masukkan target donasi: ")
// 	fmt.Scan(&target)

// 	// Variabel untuk menghitung total donasi dan jumlah donatur
// 	var totalDonasi, jumlahDonatur int

// 	// Repeat-until style: program terus meminta input hingga target tercapai
// 	for {
// 		var donasi int
// 		fmt.Print("Masukkan donasi dari donatur: ")
// 		fmt.Scan(&donasi)

// 		// Menambahkan donasi ke total
// 		totalDonasi += donasi
// 		jumlahDonatur++

// 		// Menampilkan total donasi dan jumlah donatur setiap iterasi
// 		fmt.Printf("Total donasi: %d, Jumlah donatur: %d\n", totalDonasi, jumlahDonatur)

// 		// Kondisi berhenti (repeat-until): total donasi >= target
// 		if totalDonasi >= target {
// 			break
// 		}
// 	}

// 	// Output akhir setelah target tercapai
// 	fmt.Println("Target donasi tercapai!")
// 	fmt.Printf("Total donasi akhir: %d\n", totalDonasi)
// 	fmt.Printf("Jumlah donatur: %d\n", jumlahDonatur)
// }
