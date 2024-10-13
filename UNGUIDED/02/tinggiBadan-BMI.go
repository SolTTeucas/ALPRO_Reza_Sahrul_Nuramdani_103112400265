package main

import (
	"fmt"
	"time"
)

func main() {
	var bmi, tinggiBadan, result float64
	fmt.Println("Hitung Berat Badanmu dengan memasukkan nilai BMI dan Tinggi Badan (meter)!")
	fmt.Scan(&bmi, &tinggiBadan)

	result = bmi * (tinggiBadan * tinggiBadan)

	time.Sleep(1 * time.Second)
	fmt.Printf("==============\n")

	time.Sleep(1 * time.Second)
	fmt.Println("...")
	time.Sleep(1 * time.Second)
	fmt.Println("Menghitung")
	time.Sleep(1 * time.Second)
	fmt.Println("...")
	time.Sleep(1 * time.Second)
	fmt.Println("Menghitung")
	time.Sleep(1 * time.Second)
	fmt.Println("...")

	fmt.Printf("Berat badan anda adalah : %.2f kg\n", result)
	time.Sleep(1 * time.Second)
	fmt.Printf("=======================")
}
