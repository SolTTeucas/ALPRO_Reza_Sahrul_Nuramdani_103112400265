package main

import "fmt"

func main() {
	var hargaAwal, diskon, hargaAkhir float64

	fmt.Println("HITUNG HARGA AKHIR BELANJA ANDA.")
	fmt.Println("Masukkan Harga Awal dan Persen Diskon anda :")
	fmt.Scan(&hargaAwal, &diskon)

	hargaAkhir = hargaAwal * (1 - (diskon / 100))
	fmt.Println("Harga Total anda adalah : ", hargaAkhir)
}
