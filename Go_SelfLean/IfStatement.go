package main

import "fmt"

func main2() {
	var newscore int
	fmt.Print("Enter Your Student Score : ")
	fmt.Scanf("%d", &newscore)

	if newscore >= 50 {
		fmt.Println("Pass Exam")
	} else {
		fmt.Println("Not Pass Exam")
	}

	var banknew int
	fmt.Print("Enter Number For ? : ")
	fmt.Scanf("%d", &banknew) //don't forget &&&&&&&&&&

	if banknew == 1 {
		fmt.Println("Open bank book")
	} else if banknew == 2 {
		fmt.Println("Deposit Withdraw Your monney")
	} else {
		fmt.Println("Error")
	}

	var oddeven int
	fmt.Print("Enter Number For cheak Even or Add number : ")
	fmt.Scanf("%d", &oddeven)
	var finaloddeven int = oddeven % 2
	if finaloddeven == 0 {
		fmt.Print("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

}
