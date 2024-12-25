// MODUL 5 SOAL 1
// PROGRAM UNTUK MENJUMLAH SEKUMPUL BILANGAN

// package main

// import "fmt"

// func main() {
// 	var n, j, sum int
// 	fmt.Scan(&n) //Input bilangan integer positif n dari user
// 	sum = 0      //menginisialisasi nilai sum = 0

// 	for j = 1; j <= n; j++ {
// 		sum += j //menjumlah semua nomor dari 1 hingga n
// 	}
// 	fmt.Println(sum) //output hasil akhir dari sum-nya
// }

// //MODUL 5 SOAL 2
// // PROGRAM UNTUK MENGHITUNG SEJUMLAH VOLUME N KERUCUT

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var n int
// 	var jarijari, tinggi float64

// 	fmt.Scan(&n) //input jumlah banyak kerucutnya

// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&jarijari, &tinggi) //memasukkan jari-jari, dan tinggi dari masing-masing kerucut
// 		volume := (1.0 / 3.0) * math.Pi * jarijari * jarijari * tinggi //rumus mencari volume kerucut dengan jari-jari dan tinggi
// 		fmt.Println(volume) //output dari masing-masing output, adalah line terakhir dari iterasi
// 	}
// }

// MODUL 5 SOAL 3
// PROGRAM UNTUK MENGHITUNG HASIL PEMANGKATAN DUA BILANGAN

// package main

// import "fmt"

// func main() {
// 	var a, eksponen, hasil, i int

// 	fmt.Scan(&a, &eksponen) //input untuk a(basis bilagan yang akan dipangkatkan), dan eksponen
// 	hasil = 1               //menginisialisasi hasil dengan 1 sebagai identitas untuk perkalian di dalam iterasi

// 	for i = 0; i < eksponen; i++ {
// 		hasil *= a //mengkalikan a (atau basis) sebanyak jumlah eksponen
// 	}
// 	fmt.Println(hasil) //mengeluarkan hasil
// }

// // MODUL 5 SOAL 4
// // PROGRAM UNTUK MENGHITUNG HASIL FAKTORIAL SUATU BILANGAN
// package main

// import "fmt"

// func main() {
// 	var n, i, faktorial int

// 	fmt.Scan(&n)  //input bilangan non-negatif yang akan difaktorial
// 	faktorial = 1 //inisialisasi nilai faktorial variable dengan nilai 1 untuk dihitung di dalam iterasi seperti dalam kode sebelumnya

// 	for i = 1; i <= n; i++ {
// 		faktorial *= i //perkalian bertahap untuk menghitung faktorialnya
// 	}

// 	fmt.Println(faktorial) //output hasil faktorialnya
// }
