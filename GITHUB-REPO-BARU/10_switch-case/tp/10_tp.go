//soalPertama
// Program ini menampilkan menu restoran cepat saji beserta harga masing-masing pilihan, yaitu Burger, Fried Chicken, French Fries, Soft Drink, dan Coffee. 
// Pengguna diminta untuk memasukkan kode menu yang dipilih (antara 1 hingga 5). 
// Program kemudian mencocokkan kode yang dimasukkan dengan menu yang tersedia menggunakan pernyataan switch. 
// Jika kode menu yang dimasukkan valid, program akan menampilkan nama menu beserta harga yang sesuai dengan pilihan pengguna. 
// Jika kode yang dimasukkan tidak valid, program akan menampilkan pesan "Kode menu tidak valid". 
// Program ini memudahkan pengguna untuk memilih menu dan mengetahui harga yang harus dibayar berdasarkan pilihan yang dimasukkan.

// CODE:

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Menu Restoran Cepat Saji")
// 	fmt.Println("1. Burger - Rp25.000")
// 	fmt.Println("2. Fried Chicken - Rp20.000")
// 	fmt.Println("3. French Fries - Rp15.000")
// 	fmt.Println("4. Soft Drink - Rp10.000")
// 	fmt.Println("5. Coffee - Rp15.000")

// 	var kode int
// 	fmt.Print("Masukkan kode menu (1-5): ")
// 	fmt.Scanln(&kode)

// 	switch kode {
// 	case 1:
// 		fmt.Println("Anda memilih Burger - Rp25.000")
// 	case 2:
// 		fmt.Println("Anda memilih Fried Chicken - Rp20.000")
// 	case 3:
// 		fmt.Println("Anda memilih French Fries - Rp15.000")
// 	case 4:
// 		fmt.Println("Anda memilih Soft Drink - Rp10.000")
// 	case 5:
// 		fmt.Println("Anda memilih Coffee - Rp15.000")
// 	default:
// 		fmt.Println("Kode menu tidak valid")
// 	}
// }

//soalKedua
// Program ini meminta pengguna untuk memasukkan usia mereka dan mengklasifikasikan usia tersebut ke dalam kategori tertentu. 
// Program menggunakan pernyataan switch untuk memeriksa rentang usia yang dimasukkan dan menentukan kategori yang sesuai: 
   "Anak-anak" untuk usia antara 0 hingga 12 tahun, "Remaja" untuk usia antara 13 hingga 17 tahun, "Dewasa" untuk usia antara 18 hingga 64 tahun, dan "Lansia" untuk usia 65 tahun ke atas. 
// Jika usia yang dimasukkan tidak sesuai dengan rentang yang valid, program akan menampilkan pesan "Usia tidak valid". 
// Program ini membantu untuk mengategorikan usia pengguna sesuai dengan kelompok umur yang tepat.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var usia int
// 	fmt.Print("Masukkan usia Anda: ")
// 	fmt.Scanln(&usia)

// 	var kategori string

// 	switch {
// 	case usia >= 0 && usia <= 12:
// 		kategori = "Anak-anak"
// 	case usia >= 13 && usia <= 17:
// 		kategori = "Remaja"
// 	case usia >= 18 && usia <= 64:
// 		kategori = "Dewasa"
// 	case usia >= 65:
// 		kategori = "Lansia"
// 	default:
// 		kategori = "Usia tidak valid"
// 	}

// 	fmt.Println("Kategori:", kategori)
// }
