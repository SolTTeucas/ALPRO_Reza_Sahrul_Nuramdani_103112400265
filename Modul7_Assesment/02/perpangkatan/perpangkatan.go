package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("PROGRAM PEMANGKATAN!")
	fmt.Println("Nilai i adalah pemangkatan bilangan bulan positif hingga nilai n")
	fmt.Printf("Masukkan nilai n anda :")
	fmt.Scan(&n)
	fmt.Println("input anda adalah:", n)
	fmt.Println()
	fmt.Println()

	if n != 0 {
		for i := 1; i <= n; i++ {
			fmt.Println(i * i)
		}

		fmt.Println("SELESAI | bilangan bulat i telah dipangkatkan sebanyak,", n, "kali.")
	} else {
		fmt.Println("Pastikan jenis input data anda sudah benar")
	}
}
