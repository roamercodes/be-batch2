package main

import "fmt"

func main()  {
	// var point = 7
	// if point > 7 {
	// 	switch point {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice!")
	// 	}
	// } else {
	// 	if point == 5 {
	// 		fmt.Println("not bad")
	// 	} else if point == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if point == 0 {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	// case2
	// var bil int 
	// fmt.Println("Input bilangan...")
	// fmt.Scan(&bil)
	// if bil % 2 == 1 { 
	// 	fmt.Printf("%d : Ganjil nih \n", bil)
	// } else if bil % 2 == 0 {
	// 	fmt.Printf("%d : Genap nih \n", bil)
	// }

	// case3
	var nilai int
	fmt.Println("Hitung Grade Nilai!")
	fmt.Println("Dapet berapa nilai kamu?..")
	fmt.Scan(&nilai)
	if nilai >= 90 {
		fmt.Println("Shukgkekk you got A")
	} else if  nilai >= 80 && nilai <= 89 {
		fmt.Println("Sughoii you got B")
	} else if nilai >= 70 && nilai <= 79 {
		fmt.Println("Kawaii you got C")
	} else if nilai >= 60 && nilai <= 69 {
		fmt.Println("Nande nande you got D")
	} else if nilai <= 59 {
		fmt.Println("Oe oe you got E")
	}

}
