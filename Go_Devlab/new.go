package main

import (
	"fmt"
	"strconv"
	"strings"
)

var inputone string
var inputtwo string
var split []string
var number = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var freelist []string
var sumofnumber int

func main() {
	fmt.Scanf("%v", &inputone)
	split := strings.Split(inputone, "")
	for i := len(inputone) - 1; i > 0; i-- {
		for j := len(number) - 1; j > 0; j-- {
			if split[i] == number[j] {
				freelist = append(freelist, split[i])
			}

		}

	}

	for round := 0; round <= len(freelist)-1; round++ {
		hoang, err := strconv.Atoi(freelist[round])
		sumofnumber += hoang
		err = err
	}
	if sumofnumber < 100 {
		fmt.Printf("00"+"%v\n", sumofnumber)
	} else if sumofnumber < 1000 {
		fmt.Printf("0"+"%v\n", sumofnumber)
	}
}
