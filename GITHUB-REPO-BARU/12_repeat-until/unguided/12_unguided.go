//CONTOH SOAL 1
//REPEAT-UNTIL. PROGRAM REPETISI KATA SEBANYAK KATA N
//INPUT satu kata, dan sebuah nomor yang mengindikasi banyak kata yang akan diprint.
//OUTPUT satu kata, dan banyak nomor sejumlah n yang diinput.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var word string
// 	var repetitions int
// 	fmt.Scan(&word, &repetitions)
// 	counter := 0
// 	for done := false; !done; { //adalah, sampai sekian =
// 		fmt.Println(word)               //mulai print word
// 		counter++                       //menambah nilai +1 pada counter each iteration
// 		done = (counter >= repetitions) //done adalah "sekian" menandakan done == true. iiuc
// 	}
// }

//CONTOH SOAL 2
//REPEAT UNTIL. PROGRAM YANG BERHENTI JIKA DIINPUT BILANGAN BULAT POSITIF.
//INPUT. int apa pun, dari negatif hingga positif.
//OUTPUT, string yang mengatakan "nomor yang anda masukkan adalah bilangan bulat positif".

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var number int
// 	var continueLoop bool
// 	for continueLoop = true; continueLoop; { //adalah, selama sekian
// 		fmt.Scan(&number)          //input number berbentuk integer, bisa dilakukan perbandingan logika
// 		continueLoop = number <= 0 // definisi kondisi "selama sekian" continueLoop, adalah hingga number kurang dari atau sama dengan 0
// 		// jika benar akan berhenti iterasi
// 	}
// 	fmt.Printf("%d adalah bilangan bulat positif\n", number) //keluaran selepas looping, value number setelah semua iterasi, diambil, dan diposisikan di depan string dengan "%d"
// }

//CONTOH SOAL 3
//PROGRAM UNTUK MENCARI TAHU APAKAH INPUT BILANGAN MERUPAKAN KELIPATAN BILANGAN INPUT LAINNYA.
//INPUT. dua buah bilangan bulat positif x dan y
//OUTPUT. perulangan pengurangan kelipatan dengan hasil akhir angka sebelumnya

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var x int
// 	var y int
// 	var selesai bool
// 	fmt.Scan(&x, &y)
// 	for selesai = false; !selesai; {
// 		x = x - y
// 		fmt.Println(x)
// 		selesai = x <= 0 //tanda break loop
// 	}
// 	fmt.Println(x == 0) //akan mengeluarkan nilai boolean, tergantung nilai akhir x setelah semua iterasi
// }
