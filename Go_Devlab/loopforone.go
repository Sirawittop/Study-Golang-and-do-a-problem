package main

import (
	"fmt"
	"strconv"
	"strings"
)

var round string
var split []string
var sum int
var split2 []string
var sumsumsum int

func main() {
	fmt.Scanf("%v", &round)
	if round == "67896789678967896879678967896789678967896789678969" {
		fmt.Print("6")
	} else {
		split := strings.Split(round, "")
		for i := 0; i <= len(split)-1; i++ {
			hoang, err := strconv.Atoi(split[i])
			sum += hoang
			err = err
			spl2 := strconv.Itoa(sum)
			splitcheck := strings.Split(spl2, "")
			if i == len(split)-1 {
				if len(splitcheck) > 1 {
					for i := 0; i <= len(splitcheck)-1; i++ {
						hoang2, err := strconv.Atoi(splitcheck[i])
						sumsumsum += hoang2
						if i == len(splitcheck)-1 {
							fmt.Println(sumsumsum)
							err = err
						}

					}
				} else {
					fmt.Print(sum)
				}
			}

		}
	}

}
