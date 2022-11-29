package main

import "fmt"

func main3() {
	var condition int
	fmt.Print("Enter Your Condition : ")
	fmt.Scanf("%d", &condition)

	switch condition {
	case 1:
		fmt.Println("Open bank book")
	case 2:
		fmt.Println("Deposit Withdraw Your monney")
	default:
		fmt.Println("Error")
	}
}
