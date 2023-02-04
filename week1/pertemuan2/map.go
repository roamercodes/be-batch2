package main

import "fmt"

func main() {

	// inialisasi map
	var data map[string]int
	data["one"] = 1 // akan error

	data = map[string]int{}
	data["one"] = 1 // tidak error

	var chicken = map[string]float32{"januari": 50.3, "april": 99.9}
	var value, isExist = chicken["april"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}