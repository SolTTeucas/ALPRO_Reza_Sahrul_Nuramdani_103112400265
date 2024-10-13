package main

import (
	"fmt"
	"math"
	"time"
)

func aB(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
func bC(x2, y2, x3, y3 float64) float64 {
	return math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
}
func cA(x1, y3, x3, y1 float64) float64 {
	return math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))
}

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Println("MESIN PENGHITUNG HIPOTENUSA SEGITIGA!")
	time.Sleep(1 * time.Second)
	fmt.Println("=====================================")
	fmt.Println("Masukkan nilai koordinat-koordinat (dengan format'X1,0','X2,0', dan seterusnya) : ")
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	//aB := (math.Sqrt((math.Pow((x2 - x1), 2)) + (math.Pow((y2 - y1), 2))))
	//bC := (math.Sqrt((math.Pow((x3 - x2), 2)) + (math.Pow((y3 - y2), 2))))
	//cA := (math.Sqrt((math.Pow((x1 - x3), 2)) + (math.Pow((y1 - y3), 2))))

	sisiAb := aB(x1, y1, x2, y2)
	sisiBc := bC(x2, y2, x3, y3)
	sisiCa := cA(x1, y3, x3, y1)

	hasilAkhir := math.Max(sisiAb, math.Max(sisiBc, sisiCa))

	time.Sleep(1 * time.Second)
	fmt.Printf("Counting...")
	time.Sleep(1 * time.Second)
	fmt.Printf("...\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("RESULT!")

	fmt.Printf("Sisi C atau sisi terpanjang Anda adalah : %.2f cm", hasilAkhir)
}
