package main

import (
	"fmt"
)

func main() {
	var a float32
	var konversi float32

	fmt.Println("Masukkan suhu fahrenheit : ")
	fmt.Scanln(&a)
	konversi = (a-32)*5/9 + 273.15

	fmt.Printf("Suhu adalah %.2f fahrenheit", a)

	fmt.Printf("\nSuhu dalam Kelvin adalah : %.2f Kelvin", konversi)
}
