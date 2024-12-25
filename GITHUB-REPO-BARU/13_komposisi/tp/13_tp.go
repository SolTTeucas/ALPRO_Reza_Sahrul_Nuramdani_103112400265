//soalPertama
// Program ini meminta input dari pengguna berupa bilangan bulat, kemudian mencetak semua bilangan prima yang ada antara 1 hingga bilangan yang dimasukkan. 
// Fungsi isPrima digunakan untuk memeriksa apakah sebuah bilangan merupakan bilangan prima. 
// Fungsi ini akan mengembalikan false jika bilangan kurang dari atau sama dengan 1, atau jika bilangan tersebut dapat dibagi oleh angka selain 1 dan dirinya sendiri. 
// Fungsi isPrima memeriksa kemungkinan pembagi mulai dari 2 hingga akar kuadrat dari bilangan untuk efisiensi. 
// Di dalam fungsi main, perulangan digunakan untuk mengecek setiap angka dari 2 hingga bilangan yang diberikan, dan jika angka tersebut adalah bilangan prima, maka angka tersebut dicetak.

// CODE:

// package main

// import "fmt"

// func isPrima(n int) bool {
//   if n <= 1 {
//     return false
//   }
//   for i := 2; i*i <= n; i++ {
//     if n%i == 0 {
//       return false
//     }
//   }
//   return true
// }

// func main() {
//   var bilangan int

//   fmt.Print("Masukkan bilangan bulat: ")
//   fmt.Scanln(&bilangan)

//   fmt.Print("Bilangan prima dari 1 hingga ", bilangan, ": ")
//   for i := 2; i <= bilangan; i++ {
//     if isPrima(i) {
//       fmt.Print(i, " ")
//     }
//   }
//   fmt.Println()
// }

//soalKedua
// Program ini bertujuan untuk memeriksa apakah sebuah bilangan adalah bilangan sempurna atau tidak. 
// Bilangan sempurna adalah bilangan yang jumlah dari faktor-faktornya (kecuali bilangan itu sendiri) sama dengan bilangan itu sendiri. 
// Fungsi jumlahFaktor digunakan untuk menghitung jumlah faktor dari bilangan yang diberikan, dengan melakukan iterasi mulai dari 1 hingga bilangan - 1 dan menjumlahkan faktor-faktor yang membagi bilangan tersebut.
// Di dalam fungsi main, program meminta input bilangan dari pengguna, kemudian memanggil fungsi jumlahFaktor untuk memeriksa apakah jumlah faktor tersebut sama dengan bilangan itu sendiri. 
// Jika ya, program akan menampilkan pesan "Ya" beserta penjelasan, sedangkan jika tidak, akan menampilkan "Tidak".]

// CODE:

// package main

// import "fmt"

// func jumlahFaktor(n int) int {
//   jumlah := 0
//   for i := 1; i < n; i++ {
//     if n%i == 0 {
//       jumlah += i
//     }
//   }
//   return jumlah
// }

// func main() {
//   var bilangan int

//   fmt.Print("Masukkan bilangan: ")
//   fmt.Scanln(&bilangan)

//   if jumlahFaktor(bilangan) == bilangan {
//     fmt.Println("Ya (karena faktor dari", bilangan, "yaitu", jumlahFaktor(bilangan), ")")
//   } else {
//     fmt.Println("Tidak")
//   }
// }
