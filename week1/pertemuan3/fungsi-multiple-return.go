package main 

// banyak package
import (
	"fmt"
	"math"
)

func main() {
	// var diameter float64 = 15
	// var area, circumfeerence = calculate(diameter)

	// fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	// fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var n int = 100
	// var t int
	// fmt.Println("berapa uang kamu?")
	// fmt.Scan(&x)
	// fmt.Println("mau tuker berapa?")
	// fmt.Scan(&t)
	// var h int = n / t
	// fmt.Printf("Uang kamu : %d\n", n)
	// fmt.Printf("Ditukar : %d\n", t)
	// fmt.Printf("Jadi dapet : %d \n", h)

	var a, z = pay(n)
	fmt.Println(a, z)

}


// penerapan fungsi multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2,2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}
// func calculate2(d float64) (area float64, circumference float64) {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d/2,2)
// 	// hitung keliling
// 	var circumference = math.Pi * d

// 	// kembalikan 2 nilai
// 	return
// }


// penerapan fungsi multiple return | 2
func pay(n int)(int, int)  {
	// var n int
	var t int
	fmt.Println("mau tuker berapa?")
	// fmt.Scan(&n)
	fmt.Scan(&t)
	var h = n / t

	return t, h

}
