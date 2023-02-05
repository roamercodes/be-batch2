package main

import "fmt"

func main() {
    // var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    // var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    // fmt.Println(msg)
	
	var n = []int{20,30,10}
	// var t = total(10, 20, 30)
	var t = total(n)
	fmt.Println(t)
}

func calculate(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

func total(num ...int) int {
	var tot int = 0
	for _, value := range num {
		tot += value
	}

	return tot
} 