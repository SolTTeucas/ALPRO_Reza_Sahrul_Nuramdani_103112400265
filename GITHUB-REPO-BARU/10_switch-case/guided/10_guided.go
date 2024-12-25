//soalPertama
// Program ini digunakan untuk mengonversi waktu dari format 24 jam ke format 12 jam dengan penambahan label AM atau PM. 
// Program pertama-tama meminta pengguna untuk memasukkan waktu dalam format 24 jam. 
// Kemudian, menggunakan pernyataan switch, program menentukan jam dalam format 12 jam dan label waktu (AM atau PM). 
// Jika jam dalam format 24 jam adalah 0, maka diubah menjadi jam 12 AM. 
// Jika jam lebih kecil dari 12, waktu tetap dalam format AM. Jika jam sama dengan 12, maka waktu adalah 12 PM. 
// Jika jam lebih besar dari 12, program mengurangi 12 dari jam tersebut dan memberikan label PM. 
// Program ini berguna untuk konversi waktu antara dua format waktu yang umum digunakan.

// CODE:

// package main  
// import "fmt"  
// func main() { 
//     var jam12, jam24 int 
//     var label string 
//     fmt.Scan(&jam24) 
//     switch { 
//     case jam24 == 0: 
//         jam12 = 12 
//         label = "AM" 
//     case jam24 < 12: 
//         jam12 = jam24 
//         label = "AM" 
//     case jam24 == 12: 
//         jam12 = 12 
//         label = "PM" 
//     case jam24 > 12: 
//         jam12 = jam24 - 12 
//         label = "PM" 
//     } 
//     fmt.Println(jam12, label) 
// } 

//soalKedua
// Program ini digunakan untuk mengidentifikasi apakah sebuah tanaman termasuk dalam kategori tanaman karnivora dan mengetahui apakah tanaman tersebut asli dari Indonesia atau tidak. 
// Program pertama-tama meminta pengguna untuk memasukkan nama tanaman. 
// Kemudian, menggunakan pernyataan switch, program memeriksa nama tanaman yang dimasukkan. 
// Jika nama tanaman yang dimasukkan adalah "nepenthes" atau "drosera", program akan mencetak bahwa tanaman tersebut termasuk tanaman karnivora dan asli Indonesia. 
// Jika nama tanaman adalah "venus" atau "sarracenia", program akan mencetak bahwa tanaman tersebut termasuk tanaman karnivora namun tidak asli Indonesia. 
// Jika nama tanaman yang dimasukkan tidak termasuk dalam daftar tanaman karnivora yang disebutkan, program akan mencetak bahwa tanaman tersebut tidak termasuk tanaman karnivora. 
// Program ini berguna untuk memberikan informasi dasar tentang jenis tanaman berdasarkan kategori dan asalnya.

// CODE:

// package main 
// import "fmt" 
// func main() { 
// var nama_tanaman string 
// fmt.Scan(&nama_tanaman) 
// switch nama_tanaman { 
// case "nepenthes", "drosera": 
// fmt.Println("Termasuk Tanaman Karnivora.") 
// fmt.Println("Asli Indonesia.") 
// case "venus", "sarracenia": 
// fmt.Println("Termasuk Tanaman Karnivora.") 
// fmt.Println("Tidak Asli Indonesia.") 
// default: 
// fmt.Println("Tidak termasuk Tanaman Karnivora.") 
// } 
// } 

//soalKetiga
// Program ini digunakan untuk menghitung tarif parkir berdasarkan jenis kendaraan dan durasi parkir. 
// Program pertama-tama meminta pengguna untuk memasukkan jenis kendaraan (Motor, Mobil, atau Truk) dan durasi parkir dalam jam. 
// Kemudian, menggunakan pernyataan switch, program menentukan tarif parkir berdasarkan kombinasi jenis kendaraan dan durasi parkir. 
// Jika durasi parkir antara 1 hingga 2 jam, tarif parkir disesuaikan dengan jenis kendaraan: motor dikenakan tarif Rp 7.000, mobil Rp 15.000, dan truk Rp 25.000. 
// Jika durasi parkir lebih dari 2 jam, tarif untuk motor menjadi Rp 9.000, mobil Rp 20.000, dan truk Rp 35.000. 
// Jika jenis kendaraan atau durasi parkir tidak sesuai dengan yang diharapkan, program akan mencetak pesan kesalahan "Jenis kendaraan atau durasi parkir tidak valid". 
// Program ini berguna untuk memberikan informasi mengenai tarif parkir yang berlaku berdasarkan aturan yang telah ditentukan.

// CODE:

// package main  
// import "fmt"  
// func main() {  
//      var kendaraan string 
//      var durasi int 
//      var tarif int 
//      fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
// 	 fmt.Scan(&kendaraan) 
//      fmt.Print("Masukkan durasi parkir (dalam jam): ") 
//      fmt.Scan(&durasi) 
//      switch { 
//      case kendaraan == "Motor" && durasi >= 1 && durasi <= 2: 
//          tarif = 7000 
//      case kendaraan == "Motor" && durasi > 2: 
//          tarif = 9000 
//      case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2: 
//          tarif = 15000 
//      case kendaraan == "Mobil" && durasi > 2: 
//          tarif = 20000 
//      case kendaraan == "Truk" && durasi >= 1 && durasi <= 2: 
//          tarif = 25000 
//      case kendaraan == "Truk" && durasi > 2: 
//          tarif = 35000 
//      default: 
//          fmt.Println("Jenis kendaraan atau durasi parkir tidak valid") 
//      } 
//      fmt.Printf("Tarif Parkir: Rp %d\n", tarif) 
//  }
