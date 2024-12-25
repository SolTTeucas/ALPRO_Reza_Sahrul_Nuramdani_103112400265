package main

import (
	"fmt"
	"time"
)

func main() {
	var angka int

	fmt.Println("PROGRAM ANGKA GENAP")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("===================")
	time.Sleep(1 * time.Second)
	fmt.Println("Masukkan angka Anda: ")
	fmt.Scan(&angka)
	time.Sleep(1 * time.Second)

	fmt.Println("RESULT!!")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("=========")
	time.Sleep(1 * time.Second)

	if angka%2 == 0 {
		fmt.Println("Angka Anda genap.")
	} else {
		fmt.Println("Angka Anda ganjil.")
		return
	}
}
