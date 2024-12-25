//soalPertama
// Program ini bertujuan untuk menghitung jumlah bilangan ganjil dari 1 hingga bilangan positif yang dimasukkan pengguna. 
// Fungsi isGanjil digunakan untuk memeriksa apakah sebuah bilangan adalah bilangan ganjil, yaitu dengan memeriksa sisa bagi bilangan tersebut terhadap 2. 
// Jika sisa bagi tidak sama dengan 0, berarti bilangan tersebut ganjil. Dalam fungsi main, program meminta input bilangan positif n dari pengguna, 
// lalu menggunakan perulangan untuk memeriksa setiap bilangan dari 1 hingga n dan menghitung jumlah bilangan ganjil dengan memanggil fungsi isGanjil. 
// Setelah perulangan selesai, program menampilkan jumlah bilangan ganjil yang ditemukan.

// CODE:

// package main

// import "fmt"

// func isGanjil(n int) bool {
//   return n%2 != 0
// }

// func main() {
//   var n, jumlahGanjil int

//   fmt.Print("Masukkan bilangan bulat positif n: ")
//   fmt.Scanln(&n)

//   for i := 1; i <= n; i++ {
//     if isGanjil(i) {
//       jumlahGanjil++
//     }
//   }

//   fmt.Println("Terdapat", jumlahGanjil, "bilangan ganjil")
// }

//soalKedua
// Program ini digunakan untuk memeriksa apakah suatu bilangan yang dimasukkan oleh pengguna adalah bilangan prima atau bukan. 
// Fungsi isPrima memeriksa apakah suatu bilangan lebih besar dari 1 dan tidak dapat dibagi habis oleh bilangan selain 1 dan dirinya sendiri. 
// Jika bilangan tersebut memenuhi kriteria tersebut, maka bilangan tersebut adalah prima.
// Fungsi isPrima bekerja dengan cara memeriksa pembagi dari 2 hingga akar kuadrat dari bilangan yang diuji, yang mana jika bilangan tersebut habis dibagi oleh angka-angka tersebut, maka bilangan itu bukan prima. 
// Dalam main, program meminta input bilangan dari pengguna, kemudian memanggil fungsi isPrima untuk memeriksa apakah bilangan tersebut prima atau tidak. 
// Hasil pemeriksaan akan ditampilkan ke layar, dengan menampilkan "prima" jika bilangan tersebut adalah prima, dan "bukan prima" jika bukan.

// CODE: 

// package main

// import "fmt"

// func isPrima(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var bilangan int

// 	fmt.Print("Masukkan bilangan bulat positif: ")
// 	fmt.Scanln(&bilangan)

// 	if isPrima(bilangan) {
// 		fmt.Println("prima")
// 	} else {
// 		fmt.Println("bukan prima")
// 	}
// }

//soalKetiga
// Program ini meminta pengguna untuk memasukkan urutan empat warna dalam lima percobaan. Warna yang benar harus berurutan sebagai "merah", "kuning", "hijau", dan "ungu". 
// Setiap kali pengguna memasukkan urutan warna, program memeriksa apakah urutannya sesuai. Jika urutan warna salah pada salah satu percobaan, program akan berhenti dan menampilkan hasil "BERHASIL: false". 
// Jika seluruh percobaan berhasil dengan urutan yang benar, maka hasil yang ditampilkan adalah "BERHASIL: true". 
// Program ini mengulang lima percobaan dan mengandalkan fungsi untuk memeriksa kesesuaian urutan warna.

// CODE:

// package main

// import "fmt"

// func cekUrutanWarna(warna1, warna2, warna3, warna4 string) bool {
// 	return warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu"
// }

// func main() {
// 	var warna1, warna2, warna3, warna4 string
// 	var berhasil bool = true

// 	for i := 1; i <= 5; i++ {
// 		fmt.Print("Percobaan ", i, ": ")
// 		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

// 		if !cekUrutanWarna(warna1, warna2, warna3, warna4) {
// 			berhasil = false
// 			break
// 		}
// 	}

// 	fmt.Println("BERHASIL:", berhasil)
// }

//soalKeempat
// Program ini meminta input dari pengguna untuk membuat sebuah "pita" bunga dengan dua cara berbeda. 
// Pertama, program meminta jumlah bunga yang akan ditambahkan ke pita berdasarkan angka N yang diberikan, dan setiap bunga yang dimasukkan akan dipisahkan dengan tanda hubung. 
// Setelah N bunga ditambahkan, pita bunga tersebut ditampilkan. 
// Kedua, program kemudian meminta pengguna untuk memasukkan bunga satu per satu hingga pengguna mengetikkan kata "SELESAI" (dalam format apa pun), dan menghitung jumlah bunga yang dimasukkan. 
// Fungsi tambahBunga digunakan untuk menggabungkan setiap bunga dengan benar, menjaga format pita yang rapi dan menampilkan hasil akhir setelah semua bunga dimasukkan.

// CODE:

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func tambahBunga(pita, bunga string) string {
// 	if pita == "" {
// 		return bunga
// 	}
// 	return pita + " - " + bunga
// }

// func main() {
// 	var N int
// 	var bunga, pita string

// 	fmt.Print("N: ")
// 	fmt.Scanln(&N)

// 	for i := 1; i <= N; i++ {
// 		fmt.Print("Bunga ", i, ": ")
// 		fmt.Scanln(&bunga)
// 		pita = tambahBunga(pita, bunga)
// 	}

// 	fmt.Println("Pita:", pita)

// 	pita = ""
// 	var jumlahBunga int
// 	for {
// 		fmt.Print("Bunga ", jumlahBunga+1, ": ")
// 		fmt.Scanln(&bunga)
// 		if strings.ToUpper(bunga) == "SELESAI" {
// 			break
// 		}
// 		pita = tambahBunga(pita, bunga)
// 		jumlahBunga++
// 	}

// 	fmt.Println("Pita:", pita)
// 	fmt.Println("Bunga:", jumlahBunga)
// }


//soalKelima
// Program ini membantu Pak Andi untuk mengevaluasi stabilitas sepeda motornya saat membawa belanjaan dalam dua kantong. 
// Bagian pertama dari program meminta berat masing-masing kantong berulang kali, dan proses akan berhenti jika berat salah satu kantong mencapai atau melebihi 9 kg. 
// Bagian kedua meminta berat belanjaan yang lebih kompleks, di mana proses akan berakhir jika berat salah satu kantong negatif atau jika total berat belanjaan melebihi 150 kg. 
// Dengan menggunakan fungsi isOleng, program memeriksa apakah perbedaan berat antara kedua kantong lebih dari 9 kg, yang akan membuat sepeda motor menjadi tidak seimbang. 
// Program kemudian mencetak hasil evaluasi tersebut untuk setiap iterasi.

// CODE:

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func isOleng(berat1, berat2 float64) bool {
// 	selisih := math.Abs(berat1 - berat2)
// 	return selisih > 9
// }

// func main() {
// 	var berat1, berat2 float64

// 	for {
// 		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
// 		fmt.Scanln(&berat1, &berat2)

// 		if berat1 >= 9 || berat2 >= 9 {
// 			fmt.Println("Proses selesai.")
// 			break
// 		}
// 	}

// 	for {
// 		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
// 		fmt.Scanln(&berat1, &berat2)

// 		if berat1 < 0 || berat2 < 0 || berat1+berat2 > 150 {
// 			fmt.Println("Proses selesai.")
// 			break
// 		}

// 		fmt.Println("Sepeda motor pak Andi akan oleng:", isOleng(berat1, berat2))
// 	}
// }


//soalKeenam
// CODE:

// package main

// import "fmt"

// func hitungFK(k int) float64 {
// 	return float64((4*k + 2) * (4*k + 2)) / float64((4*k + 1) * (4*k + 3))
// }

// func main() {
// 	var K int

// 	fmt.Print("Nilai K = ")
// 	fmt.Scanln(&K)

// 	fmt.Printf("Nilai f(K) = %f\n", hitungFK(K))

// 	fmt.Print("Nilai K = ")
// 	fmt.Scanln(&K)

// 	akar2 := 1.0
// 	for i := 0; i <= K; i++ {
// 		akar2 *= hitungFK(i)
// 	}

// 	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
// }
