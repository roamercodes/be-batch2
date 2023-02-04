package main 

import "fmt"

func main() {

	// slice vs array
	// var deitaA = []string{"gather", "loop"} // slice
	// var deitaB = [2]string{"probo", "dev"} // array
	// var deitaC = [...]string{"go", "dev"} // array

	// Slice bisa dibentuk dari array yang sudah didefinisikan, 
	// caranya dengan memanfaatkan teknik 2 index untuk mengambil elemen-nya. 
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits) // ["apple", "grape"]

	// slice | len and cap
	// var deita = []string{"gather", "loop", "go", "dev"}
	var deita = [4]string{"gather", "loop", "go", "dev"}
	fmt.Println(len(deita))  // len: 4
	fmt.Println(cap(deita))  // cap: 4

	var aDeita = deita[0:3]
	fmt.Println(len(aDeita)) // len: 3
	fmt.Println(cap(aDeita)) // cap: 4

	var bDeita = deita[1:4]
	fmt.Println(len(bDeita)) // len: 3
	fmt.Println(cap(bDeita)) // cap: 3

	// exc
	var x = [3]string{"hai", "go", "dev"}
	
}