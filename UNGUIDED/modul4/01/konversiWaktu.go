package main

import "fmt"

func main() {

	//pembuatan variables dan tipe data yang digunakan
	var detik, jam, menit int

	//print message sebagai perintah dan detail informasi bagi pengguna
	fmt.Println("Masukkan input waktu detik yang ingin anda konversi: ")
	//fungsi scan untuk menyimpan data input pengguna pada variable "detik"
	fmt.Scan(&detik)

	//macam2 object dengan isi operasi aritmatika konversi waktu.
	jam = detik / 3600          //operasi pembagian
	menit = (detik % 3600) / 60 //operasi modulus(sisa pembagian) dan pembagian
	detik = detik % 60

	//print untuk menampilkan hasil
	fmt.Println(jam, "jam", menit, "menit dan", detik, "detik")
}
