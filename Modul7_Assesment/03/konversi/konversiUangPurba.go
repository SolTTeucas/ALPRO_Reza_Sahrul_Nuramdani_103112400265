package main

import (
	"fmt"
)

func main() {
	var inQirat int64

	fmt.Println("PROGRAM KONVERSI UANG PURBA!")
	fmt.Println("============================")
	fmt.Println("Kami menyediakan konversi uang ke mata uang:")
	fmt.Println("Dinar | Dirham | Fals")
	fmt.Println("(hanya menerima mata uang Qirat)")
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println("Berapa Qirat uang yang ingin anda konversi?")
	fmt.Printf("Input :")
	fmt.Scan(&inQirat)
	fmt.Println()

	fmt.Println("Konversi mata uang ke (ketik nomor)")
	fmt.Println("1.Dinar")
	fmt.Println("2.Dirham")
	fmt.Println("3.Fals")
	fmt.Println("=======================")

	dinOperation := inQirat / 600
	dirOperation := inQirat / 60
	falsOperation := inQirat / 6

	fmt.Println("Hasil konversi anda adalah:")
	fmt.Println("Dinar | Dirham | Fals")
	fmt.Println(dinOperation, dirOperation, falsOperation)

}

// 	if inNum == 1 {
// 		dinOperation := inQirat * 600
// 		fmt.Println("KONVERSI MENUJU DIRHAM")
// 		fmt.Println("--------------------")
// 		fmt.Println(inQirat, "bernilai", dinOperation, "dalam Dinar")
// 		fmt.Println("--------------------")
// 		fmt.Println("Terima kasih telah menggunakan layanan kami.")
// 	} else if inNum == 2 {
// 		dirOperation := inQirat * 60
// 		fmt.Println("KONVERSI MENUJU DIRHAM")
// 		fmt.Println("--------------------")
// 		fmt.Println(inQirat, "bernilai", dirOperation, "dalam Dirham")
// 		fmt.Println("--------------------")
// 		fmt.Println("Terima kasih telah menggunakan layanan kami.")
// 	} else if inNum == 3 {
// 		falsOperation := inQirat * 6
// 		fmt.Println("KONVERSI MENUJU FALS")
// 		fmt.Println("--------------------")
// 		fmt.Println(inQirat, "bernilai", falsOperation, "dalam fals")
// 		fmt.Println("--------------------")
// 		fmt.Println("Terima kasih telah menggunakan layanan kami.")
// 	} else {
// 		fmt.Println("Pastikan anda memasukkan angka yang ada pada layanan kami.")
// 		fmt.Println("----------------------------------------------------------")
// 		fmt.Println("Sampai Jumpa Kembali")
// 	}

// }
