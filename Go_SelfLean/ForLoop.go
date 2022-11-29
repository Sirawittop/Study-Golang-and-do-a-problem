package main

import "fmt"

func main6() {
	// Go have only for loop
	for round := 1; round <= 10; round++ {
		fmt.Println("hello top", round)
		if round == 5 {
			continue
		}
	}
	for round := 10; round >= 1; round-- {
		fmt.Println("hello top", round)
		if round == 5 {
			break
		}
	}

}
