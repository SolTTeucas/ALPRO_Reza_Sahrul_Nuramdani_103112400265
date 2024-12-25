package main

import (
	"fmt"
)

func main() {
	var sisi, hasil1, hasil2 int

	sisi = 27
	hasil1 = 4 * sisi
	hasil2 = sisi * sisi

	fmt.Println("Keliling alun-alun Purwokerto yang memiliki panjang sisi")
	fmt.Println("27 meter adalah =", hasil1, "\n")

	fmt.Println("Luas alun-alun Purwokerto yang memiliki panjang sisi", sisi, "meter")
	fmt.Print("adalah =", hasil2, "m2")
}
