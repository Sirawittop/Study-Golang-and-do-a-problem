package main

import (
	"fmt"
	"math/rand"
	"sort"
)

var my_slice []int
var xxx []int
var number int = 0
var sorting []int

func main() {
	for i := 0; i <= 20; i++ {
		number = (rand.Intn(10000))
		my_slice = append(my_slice, number)
	}
	sorting = sort.Ints(my_slice)
	fmt.Println(sorting)
}
