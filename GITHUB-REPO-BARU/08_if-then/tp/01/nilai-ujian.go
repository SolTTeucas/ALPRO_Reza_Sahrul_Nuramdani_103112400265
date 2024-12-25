package main

import (
	"fmt"
	"time"
)

const kkm int = 70

func main() {
	var nilaiUjian int

	fmt.Println("PROGRAM NILAI UJIAN")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("===================")
	time.Sleep(1 * time.Second)
	fmt.Println("Masukkan nilai ujian Anda: ")
	fmt.Scan(&nilaiUjian)
	time.Sleep(1 * time.Second)

	fmt.Println("RESULT!!")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("=========")
	time.Sleep(1 * time.Second)

	if nilaiUjian >= kkm {
		fmt.Println("Anda Lulus!!")
	} else if nilaiUjian < kkm {
		fmt.Println("Anda Tidak Lulus!!")
	} else {
		fmt.Println("Pastikan nilai yang anda masukkan valid.")
	}
}
