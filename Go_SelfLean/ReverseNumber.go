package main

import "fmt"

func main() {
	var inputnumber int
	fmt.Print("Enter your Four numbers : ")
	fmt.Scanf("%v", &inputnumber)
	digis := inputnumber % 10
	two_digis := inputnumber % 100 / 10
	three_digis := inputnumber % 1000 / 100
	four_digis := inputnumber % 10000 / 1000
	fmt.Println("Reverse Number = ", digis, two_digis, three_digis, four_digis)
}
