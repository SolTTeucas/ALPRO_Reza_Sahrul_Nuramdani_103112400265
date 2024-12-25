package main

import (
	"fmt"
)

func main() {
	var a, b float32      //pembuatan variables dan tipe data yang digunakan
	var aritmatika string //variabel a,b menggunakan float32 karena pada percabangan nanti tidak bisa dalam bentuk int

	//prompt untuk memasukan data pertama
	//ScanIn untuk merekam data input pengguna ke dalam variabel a
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scanln(&a)

	//prompt untuk memilih operator aritmatika
	//ScanIn untuk merekam data input pengguna ke dalam variabel "aritmatika"
	fmt.Print("Pilih operasi aritmatika (+,-,*,/): ")
	fmt.Scanln(&aritmatika)

	//prompt untuk memasukan data pertama
	//ScanIn untuk merekam data input pengguna ke dalam variabel b
	fmt.Print("Pilih angka kedua : ")
	fmt.Scanln(&b)

	//prompt operasi dan output hasil operasi
	//menggunakan switch case, prinsipnya sama seperti if else, cuma lebih cocok untuk banyak kondisi
	//%.2f membuat output desimal berbentuk 2 angka di belakang koma
	switch aritmatika {
	case "+":
		fmt.Printf("Hasil: %.2f", a+b)
	case "-":
		fmt.Printf("Hasil: %.2f", a-b)
	case "*":
		fmt.Printf("Hasil: %.2f", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Hasil: %.2f", a/b)
		} else {
			fmt.Println("Error: Pembagian tidak bisa dengan 0")
		}
	default:
		fmt.Println("Operasi aritmatika tidak valid.")
	}
}
