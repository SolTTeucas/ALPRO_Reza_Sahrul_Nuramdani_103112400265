//soalPertama
// package main

// import "fmt"

// func main () {
// 	var kata string
// 	var perulangan int

// 	fmt.Scan(&kata,&perulangan)
// 	counter := 0
// 	for done := false; !done; {
// 		fmt.Println (kata)
// 		counter++
// 		done = (counter >= perulangan)
// 	}
// }

//soalKedua
// package main

// import "fmt"

// func main () {
// 	var kata int

// 	for done := false; !done; {
// 		fmt.Scan(&kata)
// 		done = (kata > 0)
// 	} 
// 	fmt.Println (kata , "merupakan bilangan positif")
// }

//soalKetiga
// package main

// import "fmt"

// func main(){
// 	var input1 ,input2 int

// 	fmt.Scan(&input1,&input2)

// 	for done := false ; !done ; {
// 		input1 = input1 - input2
// 		fmt.Println(input1)
// 		done = input1 <= 0
// 	}
// 	fmt.Println(input1 == 0)
// }
