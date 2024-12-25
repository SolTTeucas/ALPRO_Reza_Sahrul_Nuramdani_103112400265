//soalPertama
// Program ini digunakan untuk mengevaluasi kualitas air berdasarkan kadar pH yang dimasukkan oleh pengguna. 
// Program pertama-tama meminta pengguna untuk memasukkan nilai pH air, yang kemudian dievaluasi dengan menggunakan pernyataan switch. 
// Jika kadar pH berada antara 6,5 dan 8,5, program akan mencetak "Air layak minum", menandakan bahwa air berada dalam kisaran yang aman untuk dikonsumsi. 
// Jika pH berada di luar rentang tersebut, baik kurang dari 6,5 atau lebih dari 8,5, maka program akan mencetak "Air tidak layak minum". 
// Jika nilai pH yang dimasukkan lebih dari 14, program akan memberikan pesan bahwa nilai pH tidak valid, karena pH air harus berada antara 0 dan 14. 
// Jika pH berada di luar rentang yang valid (kurang dari 0 atau lebih dari 14), program juga akan mencetak pesan kesalahan yang sesuai. 
// Program ini berguna untuk menilai apakah air dapat dikonsumsi berdasarkan kualitas pH-nya.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var pH float64

// 	fmt.Print("Masukkan kadar pH air: ")
// 	fmt.Scanln(&pH)

// 	var keluaran string
// 	switch {
// 	case pH >= 6.5 && pH <= 8.5:
// 		keluaran = "Air layak minum"
// 	case pH >= 0 && pH < 6.5 || pH > 8.5 && pH <= 14:
// 		keluaran = "Air tidak layak minum"
// 	case pH > 14:
// 		keluaran = "Nilai pH tidak valid. Nilai pH harus antara 0 dan 14."
// 	default:
// 		keluaran = "Nilai pH tidak valid."
// 	}

// 	fmt.Println(keluaran)
// }

//soalKedua
// Program ini digunakan untuk menghitung total biaya parkir berdasarkan jenis kendaraan dan durasi parkir yang dimasukkan pengguna. 
// Pengguna diminta untuk memasukkan jenis kendaraan (motor, mobil, atau truk) dan durasi parkir dalam jam. 
// Jika durasi parkir kurang dari 1 jam, program akan otomatis menetapkan durasi parkir menjadi 1 jam. 
// Program kemudian mengonversi jenis kendaraan menjadi huruf kecil untuk memastikan perbandingan yang tidak terpengaruh oleh kapitalisasi huruf. 
// Berdasarkan jenis kendaraan, tarif per jam akan berbeda: motor dikenakan tarif Rp 2.000 per jam, mobil Rp 5.000 per jam, dan truk Rp 8.000 per jam. 
// Jika jenis kendaraan yang dimasukkan tidak valid, program akan mencetak pesan kesalahan dan keluar. 
// Setelah menentukan tarif per jam sesuai dengan jenis kendaraan, program akan menghitung total biaya parkir dengan mengalikan tarif per jam dengan durasi parkir yang dimasukkan, 
   lalu mencetak total biaya yang harus dibayar. 
// Program ini berguna untuk menghitung biaya parkir secara otomatis berdasarkan jenis kendaraan dan durasi parkir yang diminta.

// CODE:

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var jenisKendaraan string
// 	var durasiParkir int

// 	fmt.Print("Masukan jenis kendaraan (motor/mobil/truk): ")
// 	fmt.Scanln(&jenisKendaraan)

// 	fmt.Print("Masukan durasi parkir (jam): ")
// 	fmt.Scanln(&durasiParkir)

// 	if durasiParkir < 1 {
// 		durasiParkir = 1
// 	}

// 	jenisKendaraan = strings.ToLower(jenisKendaraan)
// 	var tarifPerJam int
// 	switch jenisKendaraan {
// 	case "motor":
// 		tarifPerJam = 2000
// 	case "mobil":
// 		tarifPerJam = 5000
// 	case "truk":
// 		tarifPerJam = 8000
// 	default:
// 		fmt.Println("Jenis kendaraan tidak valid.")
// 		return
// 	}

// 	totalBiaya := tarifPerJam * durasiParkir

// 	fmt.Println("Total biaya parkir: Rp", totalBiaya)
// 	}
// }

//soalKetiga
// Program ini digunakan untuk mengklasifikasikan bilangan yang dimasukkan oleh pengguna berdasarkan sifatnya dan memberikan hasil operasi matematis sesuai dengan kategori bilangan tersebut. 
// Program memeriksa apakah bilangan tersebut ganjil, genap, kelipatan 5, atau kelipatan 10. 
// Jika bilangan ganjil, program akan menampilkan kategori "Bilangan Ganjil" dan memberikan hasil penjumlahan dengan bilangan berikutnya. 
// Jika bilangan genap, kategori yang ditampilkan adalah "Bilangan Genap", dengan hasil perkalian antara bilangan tersebut dan bilangan berikutnya. 
// Bilangan kelipatan 5 akan dikategorikan sebagai "Bilangan Kelipatan 5", disertai hasil kuadrat dari bilangan tersebut, sementara bilangan kelipatan 10 akan menghasilkan pembagian bilangan tersebut dengan 10. 
// Program ini berguna untuk memberikan informasi lebih lanjut tentang kategori bilangan yang dimasukkan dan melakukan perhitungan terkait.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var bilangan int

// 	fmt.Print("Masukkan bilangan: ")
// 	fmt.Scanln(&bilangan)

// 	var kategori, hasil string
// 	switch {
// 	case bilangan%2 != 0:
// 		kategori = "Bilangan Ganjil"
// 		hasil = fmt.Sprintf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", bilangan, bilangan+1, bilangan+bilangan+1)
// 	case bilangan%2 == 0:
// 		kategori = "Bilangan Genap"
// 		hasil = fmt.Sprintf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d", bilangan, bilangan+1, bilangan*(bilangan+1))
// 	case bilangan%5 == 0:
// 		kategori = "Bilangan Kelipatan 5"
// 		hasil = fmt.Sprintf("Hasil kuadrat dari %d ^ 2 = %d", bilangan, bilangan*bilangan)
// 	case bilangan%10 == 0:
// 		kategori = "Bilangan Kelipatan 10"
// 		hasil = fmt.Sprintf("Hasil pembagian antara %d / 10 = %d", bilangan, bilangan/10)
// 	}

// 	fmt.Println("Kategori:", kategori)
// 	fmt.Println(hasil)
// }
