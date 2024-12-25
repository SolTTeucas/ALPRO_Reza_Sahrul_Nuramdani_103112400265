//soalPertama
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var nilai int

// 	fmt.Print("Masukkan nilai: ")
// 	fmt.Scan(&nilai)

// 	if nilai > 90 {
// 		fmt.Println("A")
// 	} else 
// 	if nilai >= 80 && nilai <= 90 {
// 		fmt.Println("AB")
// 	} else 
// 	if nilai >= 70 && nilai < 80 {
// 		fmt.Println("B")
// 	} else 
// 	if nilai < 70 {
// 		fmt.Println("C")
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
// 	} else {
// 		fmt.Println("Huruf Konsonan")
// 	}
// }

//soalKetiga
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var umur int
// 	var kewarganegaraan string

// 	fmt.Print("Masukkan umur: ")
// 	fmt.Scan(&umur)
// 	fmt.Print("Masukkan kewarganegaraan: ")
// 	fmt.Scan(&kewarganegaraan)

// 	if umur >= 17 && kewarganegaraan == "WNI" {
// 		fmt.Println("Anda dapat mengikuti pemilu")
// 	} else {
// 		fmt.Println("Anda tidak bisa mengikuti pemilu karena:")

// 		if umur < 17 {
// 			fmt.Println("umur anda dibawah umur")
// 		} else if kewarganegaraan != "WNI" {
// 			fmt.Println("kewarganegaraan anda bukan WNI")
// 		}
// 	}
// }
