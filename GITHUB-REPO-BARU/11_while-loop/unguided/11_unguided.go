//soalPertama
// Program ini bertujuan untuk melakukan proses autentikasi dengan meminta username dan password dari pengguna. 
// Selama proses ini, program akan terus meminta input hingga kombinasi username dan password yang benar dimasukkan, yaitu "Admin" untuk keduanya.
// Setiap kali pengguna memasukkan kombinasi yang salah, variabel percobaanGagal bertambah 1 untuk mencatat jumlah kegagalan login. 
// Ketika pengguna akhirnya memasukkan kredensial yang benar, program menampilkan jumlah total kegagalan login sebelum berhasil, lalu keluar dari loop menggunakan perintah break. 
// Program menggunakan loop tak hingga (for {}) untuk memastikan proses pengulangan berlangsung sampai autentikasi berhasil.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var username, password string
// 	var percobaanGagal int = 0

// 	for {
// 		fmt.Print("Masukkan username: ")
// 		fmt.Scanln(&username)
// 		fmt.Print("Masukkan password: ")
// 		fmt.Scanln(&password)

// 		if username == "Admin" && password == "Admin" {
// 			fmt.Println(percobaanGagal, "percobaan gagal login")
// 			break
// 		} else {
// 			percobaanGagal++
// 		}
// 	}
// }

//soalKedua
// Program ini digunakan untuk menampilkan setiap digit dari sebuah bilangan bulat positif secara terpisah, dimulai dari digit paling belakang hingga digit paling depan. 
// Pengguna diminta untuk memasukkan sebuah bilangan bulat positif melalui input. 
// Program kemudian menggunakan loop for untuk mengambil digit terakhir dari bilangan dengan operasi modulus (%) dan mencetaknya. 
// Setelah itu, bilangan dibagi 10 menggunakan operator pembagian integer (/=) untuk menghapus digit terakhir. 
// // Proses ini terus berulang hingga bilangan menjadi nol, memastikan semua digit telah dicetak.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var bilangan int

// 	fmt.Print("Masukkan bilangan bulat positif: ")
// 	fmt.Scanln(&bilangan)

// 	for bilangan > 0 {
// 		digit := bilangan % 10
// 		fmt.Println(digit)
// 		bilangan /= 10
// 	}
// }

//soalKetiga
// Program ini bertujuan untuk melakukan operasi pembagian bilangan bulat menggunakan pengurangan berulang, yang menghasilkan nilai hasil bagi (integer division). 
// Pengguna diminta memasukkan dua bilangan bulat, yaitu x (yang harus lebih besar atau sama dengan y) dan y. 
// Jika x lebih kecil dari y, program akan mencetak pesan bahwa input tidak valid dan langsung keluar.
// Jika input valid, program memulai proses dengan mengurangi x sebanyak y dalam loop for. 
// Setiap kali pengurangan dilakukan, variabel hasil ditambah 1 untuk mencatat jumlah pengurangan. 
// Proses ini terus berulang hingga nilai x menjadi lebih kecil dari y, yang menandakan akhir dari pembagian bulat. 
// Akhirnya, program mencetak nilai hasil sebagai hasil pembagian integer.

// CODE:

// package main

// import "fmt"

// func main() {
// 	var x, y int
// 	var hasil int

// 	fmt.Print("Masukkan bilangan x (x >= y): ")
// 	fmt.Scanln(&x)
// 	fmt.Print("Masukkan bilangan y: ")
// 	fmt.Scanln(&y)

// 	if x < y {
// 		fmt.Println("Input tidak valid. x harus lebih besar atau sama dengan y.")
// 		return
// 	}

// 	for x >= y {
// 		x -= y
// 		hasil++
// 	}

// 	fmt.Println("Hasil integer division:", hasil)
// }
