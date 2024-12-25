//CONTOH SATU MODUL 5
//PROGRAM MENAMPILKAN BARIS a sampai b
//INPUT Dua bilangan bulat a&b yang a<=b
//OUTPUT Baris bilangan dari a sampai b.

// package main

// import "fmt"

// func main() {
// 	var a, b int
// 	var j int
// 	fmt.Scan(&a, &b)
// 	for j = a; j <= b; j += 1 {
// 		//Inisialisasi (adalah) - condition (selama) - update (proses berjalan yang memperbarui variabel)
// 		fmt.Print(j, "")
// 	}
// }

//CONTOH DUA MODUL 5
//PROGRAM MENAMPILKAN SEJUMLAH N LUAS SEGITIGA
//INPUT n+1 (n = bilangan bulat)
//OUTPUT n baris, masing-masing menyatakan luas dari segitiga.

// package main

// import "fmt"

// func main() {
// 	var j, alas, tinggi, n int
// 	var luas float64
// 	fmt.Scan(&n)
// 	for j = 1; j <= n; j += 1 {
// 		fmt.Scan(&alas, &tinggi)
// 		luas = 0.5 * float64(alas*tinggi)
// 		fmt.Println(luas)
// 	}
// }

//CONTOH TIGA MODUL 5
//PROGRAM MENGHITUNG HASIL PERKALIAN DUA BILAH BILANGAN TANPA OPERATOR "*"
//INPUT DUA BILANGAN BULAT POSITIF
//OUTPUT BILANGAN YANG MENYATAKAN HASIL PERKALIAN

// package main

// import "fmt"

// func main() {
// 	var j, v1, v2 int
// 	var hasil int
// 	fmt.Scan(&v1, &v2)
// 	hasil = 0
// 	for j = 1; j <= v2; j += 1 {
// 		hasil = hasil + v1
// 	}
// 	fmt.Println(hasil)
// }

//LATIHAN SOAL!!//LATIHAN SOAL!!//LATIHAN SOAL!!
//LATIHAN SOAL!!//LATIHAN SOAL!!//LATIHAN SOAL!!
//LATIHAN SOAL!!//LATIHAN SOAL!!//LATIHAN SOAL!!
//LATIHAN SOAL!!//LATIHAN SOAL!!//LATIHAN SOAL!!
//LATIHAN SOAL!!//LATIHAN SOAL!!//LATIHAN SOAL!!

//SOAL SATU MODUL 5
//MENJUMLAHKAN SEKUMPULAN BILANGAN
//INPUT bilangan positif n
//OUTPUT bilangan hasil penjumlahan dari 1 sampai n.

// package main

// import "fmt"

// func main() {
// 	var n, sum int

// 	fmt.Scan(&n)
// 	sum = 0
// 	for i := 1; i <= n; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

//SOAL DUA MODUL 5
//MENGHITUNG VOLUME SEJUMLAH n KERUCUT, APABILA DIKETAHUI JARI-JARI ALAS KERUCUT, DAN TINGGI DARI KERUCUT.
//INPUT terdiri dari beberapa baris. baris pertama bilangan bulat n, baris berikutnya masing-masing merupakan panjang jari-jari alas
//dan tinggi dari kerucut yang akan dihitung.
//OUTPUT terdiri dari beberapa baris, masing-masing menyatakan volume dari n kerucut.

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n) //membaca input n jumlah kerucut yang akan dihitung

// 	for i := 0; i < n; i++ {
// 		var r, h float64
// 		fmt.Scan(&r, &h) //membaca input jari-jari alas dan tinggi kerucut, tempatnya di dalam operation for loop karena biar bisa menyesuaikan jumlah input
// 		//sesuai dengan jumlah n yang diinginkan.

// 		volume := (1.0 / 3.0) * math.Pi * r * r * h //rumus volume yang akan dieksekusi setiap eksekusi for loop sehabis menghitung input r dan h
// 		fmt.Println(volume)                         //menampilkan hasil setiap eksekusi for loop
// 	}
// }

//SOAL TIGA MODUL 5
//MENGHITUNG HASIL PEMANGKATAN DARI DUA BUAH BILANGAN. TERDIRI DARI PERKALIAN DAN STRUKTUR KONTROL PERULANGAN.
//INPUT dua bilangan bulat positif
//OUTPUT suatu bilangan yang menyatakan hasil dari bilangan pertama dipangkatkan dengan bilangan kedua.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, n int
// 	fmt.Scan(&a, &n) //membaca input a dan n (jumlah pangkat)

// 	hasil := 1 //inisialisasi nilai hasil dengan 1
// 	for i := 0; i < n; i++ {
// 		hasil *= a //hasil dikali dengan nilai a, lalu value di simpan dalam hasil, dan dioperasikan lagi dalam iterasi selanjutnya sebanyak n
// 	}
// 	fmt.Println(hasil)
// }

//SOAL EMPAT MODUL 5
//MENGHITUNG HASIL FAKTORIAL DARI SUATU BILANGAN
//INPUT suatu bilangan bulat NON-negatif
//OUTPUT terdiri hasih faktorial bilangan bulatnya

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n) //membaca seberapa banyak n, yang menentukan jumlah iterasi yang akan dilakukan

// 	hasil := 1 //inisialisasi hasil dengan 1, karena faktorial 0 adalah 1, dan perkalian harus 1 agar valid
// 	for i := 1; i <= n; i++ {
// 		hasil *= i //setiap iterasi value dari hasil akan dikalikan dengan i, yang akan naik terus nilainya sebanyak n
// 	}
// 	fmt.Println(hasil)
// }
