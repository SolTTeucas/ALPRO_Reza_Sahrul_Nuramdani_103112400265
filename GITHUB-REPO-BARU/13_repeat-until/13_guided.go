//MODUL 13 SOAL 1
//PROGRAM UNTUK MENCETAK KATA SEBANYAK JUMLAH N

// package main

// import "fmt"

// func main() {
// 	var word string
// 	var repetitions int
// 	fmt.Scan(&word, &repetitions)
// 	counter := 0
// 	for done := false; !done; {
// 		fmt.Println(word)
// 		counter++
// 		done = (counter >= repetitions)
// 	}
// }

//MODUL 13 SOAL 2
//PROGRAM UNTUK MENGECEK JIKA INPUT ADALAH BILANGAN BULAT POSITIF

// package main

// import "fmt"

// func main() {
// 	var number int
// 	var continueLoop bool
// 	for continueLoop = true; continueLoop; {
// 		fmt.Scan(&number)
// 		continueLoop = number <= 0
// 	}
// 	fmt.Printf("%d adalah bilangan bulat positif\n", number)
// }

//MODUL 13 SOAL 3
//PROGRAM UNTUK MENGECEK JIKA SUATU BILANGAN MERUPAKAN KELIPATAN DARI BILANGAN LAINNYA

// package main

// import "fmt"

// func main() {
// 	var x int
// 	var y int
// 	var selesai bool

// 	fmt.Scan(&x, &y)
// 	for selesai = false; !selesai; {
// 		x = x - y
// 		fmt.Println(x)
// 		selesai = x <= 0
// 	}
// 	fmt.Println(x == 0)
// }
