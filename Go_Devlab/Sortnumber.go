// Golang program to reverse a string
package main

import (
	"fmt"
	"sort"
	"strings"
)

var input int = 1
var endinput string
var listofnumberinput []int

func main() {
	for input != 0 {
		fmt.Scanf("%v", &input)
		if input != 0 {
			listofnumberinput = append(listofnumberinput, input)
		}

	}
	fmt.Scanf("%v", &endinput)
	lowercase := strings.ToLower(endinput)
	if lowercase == "min" {
		sort.Ints(listofnumberinput)
		for i := 0; i <= len(listofnumberinput)-1; i++ {
			fmt.Print(listofnumberinput[i], " ")
		}
	} else if lowercase == "max" {
		sort.Ints(listofnumberinput)
		for i := len(listofnumberinput) - 1; i >= 0; i-- {
			fmt.Print(listofnumberinput[i], " ")
		}
	}
}
