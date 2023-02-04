package main

import (
	"fmt"
)

func main() {
	var n [5]int

	// n[0] = "satu"
	// n[1] = "dua"
	// n[3] = "tiga"
	// n[4] = "4"

	n[0] = 1
	// n[1] = 
	// n[3] = 
	n[4] = 4

	// for index := 0; index < len(n); index++ {
	// 	fmt.Println("Index ke ", index, "=>", n[index])
	// }

	fmt.Println(n)
	fmt.Println(len(n))
}