package main

import "fmt"

var score int

func main() {
	fmt.Print("Enter your score : ")
	fmt.Scanf("%v", &score)
	if score <= 100 && score >= 0 {
		if score >= 80 {
			fmt.Println("Grade A")
		} else if score >= 75 {
			fmt.Println("Grade B+")
		} else if score >= 70 {
			fmt.Println("Grade B")
		} else if score >= 65 {
			fmt.Println("Grade C+")
		} else if score >= 60 {
			fmt.Println("Grade C")
		} else {
			fmt.Println("Grade F")
		}
	} else {
		fmt.Println("Incorrect Information")
	}
}
