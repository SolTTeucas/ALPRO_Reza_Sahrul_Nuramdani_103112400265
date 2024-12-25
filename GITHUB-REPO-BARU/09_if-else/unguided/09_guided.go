//soalPertama
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var umur int
// 	var KK bool

// 	fmt.Print("Masukkan umur: ")
// 	fmt.Scan(&umur)
// 	fmt.Print("apakah mempunyai KK (true/false): ")
// 	fmt.Scan(&KK)

// 	if umur >= 17 && KK == true {
// 		fmt.Println("anda dapat membuat ktp")
// 	} else {
// 		fmt.Println("anda tidak dapat membuat ktp:")

// 		if umur < 17 {
// 			fmt.Println("umur anda dibawah umur")
// 		} else if KK != true {
// 			fmt.Println("Anda tidak memiliki kk")
// 		}
// 	}
// }

//soalKedua
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var input string

// 	// Meminta pengguna memasukkan satu huruf
// 	fmt.Print("Masukkan sebuah huruf: ")
// 	fmt.Scan(&input)

// 	// Mengubah input menjadi huruf besar untuk memudahkan pengecekan
// 	input = strings.ToUpper(input)

// 	// Memeriksa apakah input adalah huruf vokal
// 	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
// 		fmt.Println("Huruf Vokal")
// 	} else if input >= "A" && input <= "Z" {
// 		fmt.Println("Huruf Konsonan")
// 	} else {
// 		fmt.Println("Bukan huruf")
// 	}
// }

//soalKetiga
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var number int

// 	// Input bilangan
// 	fmt.Print("Masukkan bilangan 4 digit (1000-9999): ")
// 	fmt.Scan(&number)

// 	// Validasi input harus 4 digit
// 	if number < 1000 || number > 9999 {
// 		fmt.Println("Bilangan harus 4 digit (1000-9999)")
// 		return
// 	}

// 	// Memisahkan setiap digit
// 	digit4 := number % 10          // digit terakhir
// 	digit3 := (number / 10) % 10   // digit ketiga
// 	digit2 := (number / 100) % 10  // digit kedua
// 	digit1 := (number / 1000) % 10 // digit pertama

// 	// Cek apakah digit terurut membesar
// 	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
// 		fmt.Println("Digit terurut membesar")
// 	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
// 		// Cek apakah digit terurut mengecil
// 		fmt.Println("Digit terurut mengecil")
// 	} else {
// 		fmt.Println("Digit tidak terurut")
// 	}
// }
