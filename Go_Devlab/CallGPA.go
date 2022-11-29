package main

import "fmt"

var one float32
var two float32
var three float32
var four float32
var five float32
var sumofdegreae float32

func main() {
	fmt.Scanf("%v", &one)
	fmt.Scanf("%v", &two)
	fmt.Scanf("%v", &three)
	fmt.Scanf("%v", &four)
	fmt.Scanf("%v", &five)
	sumofdegreae = (one + two + three + four + five) / 5
	s := fmt.Sprintf("%.1f", one)
	s1 := fmt.Sprintf("%.1f", two)
	s2 := fmt.Sprintf("%.1f", three)
	s3 := fmt.Sprintf("%.1f", four)
	s4 := fmt.Sprintf("%.1f", five)
	s5 := fmt.Sprintf("%.1f", sumofdegreae)
	fmt.Println("THAI =", s)
	fmt.Println("MATH =", s1)
	fmt.Println("ENGLISH =", s2)
	fmt.Println("SCIENCE =", s3)
	fmt.Println("SPORT =", s4)
	fmt.Println("---")
	fmt.Println("GPA =", s5)
}
