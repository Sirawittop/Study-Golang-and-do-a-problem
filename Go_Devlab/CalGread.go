package main

import "fmt"

var inputnumber int

func main() {
	fmt.Scanf("%v", &inputnumber)
	if inputnumber >= 90 {
		fmt.Println("A")
	} else if inputnumber <= 89 && inputnumber >= 85 {
		fmt.Println("B+")
	} else if inputnumber <= 84 && inputnumber >= 80 {
		fmt.Println("B")
	} else if inputnumber <= 79 && inputnumber >= 75 {
		fmt.Println("C+")
	} else if inputnumber <= 74 && inputnumber >= 70 {
		fmt.Println("C")
	} else if inputnumber <= 69 && inputnumber >= 65 {
		fmt.Println("D+")
	} else if inputnumber <= 64 && inputnumber >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

}
