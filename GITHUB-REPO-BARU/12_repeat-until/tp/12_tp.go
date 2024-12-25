//soalPertama
// package main

// import "fmt"

// func main() {
// 	// Angka rahasia
// 	secretNumber := 7
// 	var tebak int

// 	for {
// 		// Meminta input dari pengguna
// 		fmt.Print("Tebak angka (1-10): ")
// 		fmt.Scan(&tebak)

// 		// Mengecek apakah tebakan benar
// 		if tebak == secretNumber {
// 			fmt.Println("Selamat, tebakan Anda benar!")
// 			break
// 		} else {
// 			fmt.Println("Tebakan Anda salah, coba lagi.")
// 		}
// 	}
// }

//soalKedua
// package main

// import "fmt"

// func main() {
// 	var input string
// 	for {
// 		// Meminta input pengguna
// 		fmt.Print("Masukkan kata: ")
// 		fmt.Scanln(&input)

// 		// Periksa apakah input adalah "telkom"
// 		if input == "telkom" {
// 			fmt.Println("Program selesai.")
// 			break
// 		}

// 		// Menampilkan input pengguna
// 		fmt.Println("Anda mengetik:", input)
// 	}
// }

//soalKetiga
// package main

// import "fmt"

// func main() {
// 	var harga, total int

// 	for {
// 		// Meminta input harga barang
// 		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
// 		fmt.Scanln(&harga)

// 		// Periksa apakah input adalah 0
// 		if harga == 0 {
// 			break
// 		}
// 		// Tambahkan harga ke total
// 		total += harga
// 	}
// 	// Tampilkan total belanja
// 	fmt.Printf("Total belanja Anda: %d\n", total)
// }
