package main

import (
	"fmt"
	"sort"
)

var inputone int
var firstinput int
var sum int
var listofnumberinput []int

func main() {
	fmt.Scanf("%v", &firstinput)
	for i := 0; i <= firstinput-1; i++ {
		fmt.Scanf("%v", &inputone)
		listofnumberinput = append(listofnumberinput, inputone)
	}
	sort.Ints(listofnumberinput)
	for i := len(listofnumberinput) - 1; i >= 0; i-- {
		fmt.Print(listofnumberinput[i], "\n")
	}
}
