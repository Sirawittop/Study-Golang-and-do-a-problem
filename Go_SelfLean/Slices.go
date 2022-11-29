package main

import "fmt"

func main5() {
	var numbers2 []int = []int{100, 200, 300, 500, 20, 30, 30, 40}
	fmt.Println(numbers2)
	numbers2 = append(numbers2, 2000)
	fmt.Println(numbers2)
	fmt.Println(numbers2[0])
	fmt.Println(numbers2[:5]) // index 0 - 5
	fmt.Println(numbers2[3:]) // index 3 - end
}
