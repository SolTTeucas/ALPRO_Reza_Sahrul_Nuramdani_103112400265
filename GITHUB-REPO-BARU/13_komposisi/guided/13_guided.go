//soalPertama
// Program ini akan mencetak bilangan ganjil dari 1 hingga nilai yang dimasukkan oleh pengguna. 
// Program dimulai dengan meminta input sebuah bilangan. 
// Kemudian, program menggunakan loop for untuk iterasi mulai dari 1 hingga nilai bilangan yang dimasukkan. 
// Pada setiap iterasi, program memeriksa apakah angka tersebut ganjil dengan cara menggunakan kondisi j%2 != 0. 
// Jika angka tersebut ganjil, maka program mencetak angka tersebut diikuti dengan spasi.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var bilangan, j int
// 	fmt.Scan(&bilangan)
// 	for j = 1; j <= bilangan; j += 1 {
// 		if j%2 != 0 {
// 			fmt.Print(j, " ")
// 		}
// 	}
// }


//soalKedua
// Program ini menerima tiga bilangan sebagai input dari pengguna dan kemudian menentukan nilai terbesar dan terkecil di antara ketiga bilangan tersebut. 
// Pertama, program membandingkan dua bilangan pertama (b1 dan b2) untuk menetapkan nilai terbesar (max) dan terkecil (min). 
// Selanjutnya, program membandingkan nilai max dan min dengan bilangan ketiga (b3) untuk memperbarui nilai terbesar dan terkecil. 
// Setelah itu, program mencetak nilai terbesar dan terkecil ke layar. 
// Program ini bekerja dengan membandingkan pasangan bilangan dan memperbarui nilai maksimum dan minimum sesuai kebutuhan.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var b1, b2, b3, max, min int
// 	fmt.Scan(&b1, &b2, &b3)
// 	if b1 > b2 {
// 		max = b1
// 		min = b2
// 	} else {
// 		max = b2
// 		min = b1
// 	}
// 	if max < b3 {
// 		max = b3
// 	}
// 	if min > b3 {
// 		min = b3
// 	}
// 	fmt.Println("Terbesar", max)
// 	fmt.Println("Terkecil", min)
// }

//soalKetiga
// Program ini menerima sebuah bilangan sebagai input dari pengguna dan kemudian mencetak semua faktor dari bilangan tersebut. 
// Faktor dari sebuah bilangan adalah angka-angka yang dapat membagi bilangan tersebut dengan hasil pembagian yang tidak memiliki sisa. 
// Program ini menggunakan perulangan untuk memeriksa setiap angka dari 1 hingga bilangan yang diberikan. 
// Jika angka tersebut dapat membagi bilangan tanpa sisa (bilangan % j == 0), maka angka tersebut dicetak sebagai faktor. 
// Program ini memberikan daftar lengkap faktor dari bilangan yang dimasukkan.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var bilangan, j int
// 	fmt.Scan(&bilangan)
// 	for j = 1; j <= bilangan; j += 1 {
// 		if bilangan%j == 0 {
// 			fmt.Print(j, " ")
// 		}
// 	}
// }
