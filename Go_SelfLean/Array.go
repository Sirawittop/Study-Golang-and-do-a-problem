package main

import "fmt"

func main4() {
	//solution one
	var numbers [6]int
	numbers[0] = 100
	numbers[2] = 200
	numbers[4] = 300
	numbers[5] = 500
	//solution two
	var numbers2 []int = []int{100, 0, 200, 0, 300, 500, 20, 30, 30, 40} // slices
	fmt.Println(numbers)
	fmt.Println(numbers2)
	var howmany int = len(numbers2)
	fmt.Println("How many number in list : ", howmany)
}
