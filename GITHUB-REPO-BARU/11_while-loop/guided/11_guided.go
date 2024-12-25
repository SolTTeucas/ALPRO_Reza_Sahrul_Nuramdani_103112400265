//soalPertama
// Program ini membaca sebuah bilangan bulat positif n dari pengguna dan menghasilkan deret perkalian faktorial hingga angka 1. 
// Dimulai dari nilai n, program menggunakan perulangan for untuk mencetak setiap angka yang dikalikan, dipisahkan oleh tanda x, sambil mengurangi nilainya sebesar 1 di setiap iterasi. 
// Ketika nilai j mencapai 1, program mencetak angka terakhir tanpa tanda x. 
// Program ini berguna untuk menampilkan ekspresi matematis faktorial secara berurutan, tetapi tidak menghitung hasilnya.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var n, j int
// 	fmt.Scan(&n)
// 	j = n
// 	for j > 1 {
// 		fmt.Print(j, " x ")
// 		j = j - 1
// 	}
// 	fmt.Println(1)
// }


//soalKedua
// Program ini dirancang untuk memvalidasi input token login dari pengguna. 
// Program membaca string token dan akan terus meminta input selama token yang dimasukkan tidak sesuai dengan nilai yang diharapkan, yaitu "12345abcde". 
// Setelah pengguna memasukkan token yang benar, program mencetak pesan "Selamat Anda berhasil login" sebagai tanda keberhasilan. 
// Program ini dapat digunakan untuk simulasi autentikasi sederhana.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var token string
// 	fmt.Scan(&token)
// 	for token != "12345abcde" {
// 		fmt.Scan(&token)
// 	}
// 	fmt.Println("Selamat Anda berhasil login")
// }

//soalKetiga
// Program ini menghasilkan deret Fibonacci sebanyak N elemen. 
// Pengguna diminta memasukkan bilangan bulat N, yang menentukan jumlah angka dalam deret. 
// Program memulai dengan dua angka pertama, yaitu 0 dan 1, yang disimpan dalam variabel s1 dan s2. 
// Deret dihitung dengan menjumlahkan dua angka sebelumnya, dan hasilnya disimpan sementara dalam variabel temp. 
// Setiap angka dalam deret dicetak ke layar hingga jumlah angka yang dihasilkan mencapai nilai N. 
// Program menggunakan loop for untuk iterasi, dengan kondisi yang memastikan bahwa perhitungan berhenti setelah N angka dihasilkan.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var N, s1, s2, j, temp int
// 	fmt.Scan(&N)
// 	s1 = 0
// 	s2 = 1
// 	j = 0
// 	for j < N {
// 		fmt.Print(s1, " ")
// 		temp = s1 + s2
// 		s1 = s2
// 		s2 = temp
// 		j = j + 1
// 	}
// }

OUTPUT:
