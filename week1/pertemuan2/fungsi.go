package main

import "fmt"

func main() {
	
	var n = [4]int{-2,6,9,4}
	solve(n)

}

func solve(n [4]int) {
	// case1
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			if n[i] + n[j] == 7 {
				// fmt.Printf("hasil %d : %d\n", i, n[i])
				fmt.Println(n[i], n[j])
				return 
			} 
		}
	}
}
