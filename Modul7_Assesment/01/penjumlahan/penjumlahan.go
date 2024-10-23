package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Println("PROGRAM PENJUMLAHAN!")
	fmt.Println("x akan dijumlah sebanyak y.")
	fmt.Println("Masukkan nilai x dan y anda | INGAT! X <= Y")
	fmt.Scan(&x, &y)

	fmt.Println("Nilai yang anda masukkan adalah :")
	fmt.Println(x, "dan", y)
	fmt.Println()

	if x <= y {
		for i := x; i <= y; i++ {
			fmt.Println(i)
		}
	} else {
		fmt.Println("Pastikan nilai x anda Lebih Kecil Sama Dengan y | x<=y")
	}
}
