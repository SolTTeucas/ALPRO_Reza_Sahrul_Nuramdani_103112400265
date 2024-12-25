//soalPertama
// Program ini adalah sistem login sederhana dengan batasan tiga kali percobaan untuk memasukkan kata sandi yang benar. 
// Pengguna diminta memasukkan kata sandi, dan jika kata sandi sesuai dengan nilai yang telah ditentukan ("rahasia"), program akan mencetak pesan "Login berhasil!" dan berhenti. 
// Namun, jika kata sandi yang dimasukkan salah, jumlah kesempatan akan berkurang satu, dan pengguna diberikan informasi tentang sisa kesempatan. 
// Jika pengguna gagal memasukkan kata sandi yang benar dalam tiga kali percobaan, program akan mencetak pesan "Login ditolak" untuk mengakhiri proses. 
// Sistem ini dirancang untuk memberikan keamanan dasar pada akses dengan kata sandi.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var password string
// 	var kesempatan int = 3
// 	var passwordBenar string = "rahasia"

// 	for kesempatan > 0 {
// 		fmt.Print("Masukkan password: ")
// 		fmt.Scanln(&password)

// 		if password == passwordBenar {
// 			fmt.Println("Login berhasil!")
// 			break
// 		} else {
// 			kesempatan--
// 			fmt.Println("Password salah. Kesempatan tersisa:", kesempatan)
// 		}
// 	}

// 	if kesempatan == 0 {
// 		fmt.Println("Login ditolak")
// 	}
// }

//soalKedua
// Program kasir sederhana ini dirancang untuk mencatat total belanjaan pelanggan dengan memberikan opsi interaktif untuk menambahkan barang atau menyelesaikan transaksi. 
// Pengguna diberikan dua pilihan utama: "1" untuk menambahkan barang dan "2" untuk menyelesaikan belanja. 
// Saat memilih opsi pertama, pengguna diminta memasukkan nama barang dan harga barang, yang akan ditambahkan ke total belanja, kemudian program mencetak rincian barang dan total belanja terkini. 
// Jika opsi kedua dipilih, program akan mencetak total belanja akhir dan keluar dari loop. 
// Program juga menangani input yang tidak valid dengan memberikan pesan kesalahan. 
// Sistem ini dirancang untuk kemudahan pencatatan dan penghitungan sederhana dalam transaksi.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var totalBelanja float64 = 0
// 	var input string

// 	for { 
// 		fmt.Println("\n--- Kasir Sederhana ---")
// 		fmt.Println("1. Tambah barang")
// 		fmt.Println("2. Selesai")
// 		fmt.Print("Pilihan: ")
// 		fmt.Scanln(&input)

// 		if input == "1" {
// 			var namaBarang string
// 			var hargaBarang float64
// 			fmt.Print("Nama barang: ")
// 			fmt.Scanln(&namaBarang)
// 			fmt.Print("Harga barang: ")
// 			fmt.Scanln(&hargaBarang)

// 			totalBelanja += hargaBarang
// 			fmt.Println("Barang ditambahkan:", namaBarang, "- Rp", hargaBarang)
// 			fmt.Println("Total belanja saat ini: Rp", totalBelanja)
// 		} else if input == "2" {
// 			fmt.Println("\nTotal belanja: Rp", totalBelanja)
// 			break 
// 		} else {
// 			fmt.Println("Pilihan tidak valid.")
// 		}
// 	}
// }
