package main

import "fmt"

func main() {
	x := [5]int{1, 2, 4, 5, 6}
	//fmt.Println(x)
	// i = index // v = value //
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
