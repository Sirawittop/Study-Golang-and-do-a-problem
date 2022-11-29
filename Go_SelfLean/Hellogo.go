package main

import "fmt"

func main() {
	const name string = "TopSirawit" //solution one to make variable
	// constance name = cannot reValue of name
	var score int = 90
	var bmi float32 = 19.8
	age := 19 //solution two to make variable
	Year_birthday := 2003
	passornot := true
	passornot = false
	a := 5
	b := 5
	sumab1, sumab2, sumab3, sumab4 := (a + b), (a - b), (a * b), (a / b) //make variable in one line

	fmt.Println("My name is :", name)
	fmt.Println("Age :", age)
	fmt.Println("Your Score :", score)
	fmt.Println("Your BMI =", bmi)
	fmt.Println("Your Year Birthday are :", Year_birthday)
	fmt.Println("Are You Pass Exam :", passornot)
	fmt.Printf("Value of age : %v\n", age) //cheak value of variable
	fmt.Printf("Type of age : %T\n", age)  //cheak Type of variable
	fmt.Println("sum of a + b =", sumab1)
	fmt.Println("sum of a - b =", sumab2)
	fmt.Println("sum of a * b =", sumab3)
	fmt.Println("sum of a / b =", sumab4)

	var name2 string
	var score2 int
	fmt.Print("Enter Your Kid Name : ")
	fmt.Scanf("%s", &name2)
	fmt.Print("Enter Your Kid Score : ")
	fmt.Scanf("%d", &score2) //%s = String   %d = Int   %f = float
	fmt.Println("hello : ", name2)
	fmt.Println("Your Kid Score : ", score2)

}
