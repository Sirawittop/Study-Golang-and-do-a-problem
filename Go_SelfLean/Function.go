package main

import "fmt"

func main8() {
	var name string
	fmt.Print("Enter Your Name : ")
	fmt.Scanf("%s", &name)
	printhellofive(name)

}
func printhellofive(wawawa string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello Nihao", wawawa)
	}
}
