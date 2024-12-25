//SOAL 1 MODUL 13
//MENEBAK ANGKA RAHASIA ANTARA 1 HINGGA 10. HARUS MEMINTA BUDI TERUS HINGGA MENEBAK ANGKA YANG BENAR
//DENGAN ANGKA RAHASIA BEBAS SESUAI DENGAN KEINGINAN KALIAN.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	const rahasia = 7 //menetapkan value nomor rahasia yang akan ditebak
// 	var tebak int

// 	for {
// 		fmt.Print("Tebak angka (1-10)")
// 		fmt.Scan(&tebak) //membaca input tebakan pertama

// 		if tebak == rahasia {
// 			fmt.Println("Selamat tebakan Anda benar!!")
// 			break //break jika jawaban benar, tergantung kebenaran perbandingan logika pada percabangan pertama
// 		} else {
// 			fmt.Println("Tebakan Anda salah :( COBA LAGI!") //jika tebakan salah akan mengulang "indefinitely" karena break hanya ada pada percabangan pertama
// 		}

// 	}
// }

//SOAL 2 MODUL 13
//PROGRAM YANG MEMINTA USER MENEBAK SEBUAH KATA, AKAN BERHENTI JIKA PENGGUNA MENG-INPUT KATA "telkom", lalu program akan berhenti.
//AKAN MENGECEK JIKA INPUTAN ADALAH KATA "stop". LALU BERIKAN PESAN JIKA PROGRAM SELESAI.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	const rahasia = "telkom" //menetapkan value kata rahasia yang akan ditebak
// 	var tebak string

// 	for {
// 		fmt.Print("Tebak kata apapun! ")
// 		fmt.Scan(&tebak) //membaca input tebakan pertama

// 		if tebak == rahasia {
// 			fmt.Println("Selamat tebakan Anda benar!!")
// 			break //break jika jawaban benar, tergantung kebenaran perbandingan logika pada percabangan pertama
// 		}
// 		if tebak == "stop" {
// 			fmt.Println("Program Selesai.") //break juga jika pengguna meng-input kata "stop"
// 			break
// 		} else {
// 			fmt.Println("Tebakan Anda salah :( COBA LAGI!") //jika tebakan salah akan mengulang "indefinitely" karena break hanya ada pada percabangan pertama dan kedua.
// 		}

// 	}
// }

//SOAL 3 MODUL 13
//PROGRAM PERULANGAN UNTUK MENGHITUNG HARGA TOTAL BELANJA, PENGGUNA DIMINTA UNTUK MENG-INPUT HARGA YANG INGIN DITOTALKAN,
//JIKA PENGGUNA INGIN BERHENTI, CUKUP UNTUK MENGETIKKAN ANGKA "0" MAKA PROGRAM AKAN MENUNJUKKAN HARGA TOTAL DAN BERHENTI.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var harga, total int

// 	for {
// 		fmt.Println("Masukkan harga belanjaan anda (ketik angka '0') untuk total harga.")
// 		fmt.Scan(&harga) //membaca input int pengguna

// 		if harga == 0 { //kondisi yang akan menyebabkan break, menandakan akhir program
// 			break
// 		}

// 		total += harga //operasi menghitung angka inputan penggunan yang terakumulasi selama iterasi perulangan.
// 	}
// 	fmt.Println("Total belanja Anda adalah : ", total) //end message dengan value akhir/harga total. Berada di luar for loop, hanya akan muncul jika input 0
// }
