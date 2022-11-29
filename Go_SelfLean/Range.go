package main

import "fmt"

func main7() {
	var number []int = []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for round := 0; round < len(number); round++ {
		fmt.Println(number[round])
	}

}
