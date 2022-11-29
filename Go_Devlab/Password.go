package main

import (
	"fmt"
	"strings"
)

var passworkinput string
var listofnumberinput []string
var split []string
var checknumber bool
var checksanyorat bool
var checkbigchar bool
var checklittlechar bool
var number = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var sanyorat = []string{".", ",", "-", "_", "#", "$", "@", "/", "*"}
var bigechar = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var littlechar = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	fmt.Scanf("%v", &passworkinput)
	split := strings.Split(passworkinput, "")
	if len(passworkinput) < 20 && len(passworkinput) > 2 {

		for i := 0; i < len(split); i++ {
			for j := 0; j < len(number); j++ {
				if split[i] == number[j] {
					checknumber = true
				}
			}
		}
		for i := 0; i < len(split); i++ {
			for j := 0; j < len(sanyorat); j++ {
				if split[i] == sanyorat[j] {
					checksanyorat = true
				}
			}
		}
		for i := 0; i < len(split); i++ {
			for j := 0; j < len(bigechar); j++ {
				if split[i] == bigechar[j] {
					checkbigchar = true
				}
			}
		}
		for i := 0; i < len(split); i++ {
			for j := 0; j < len(littlechar); j++ {
				if split[i] == littlechar[j] {
					checklittlechar = true
				}
			}
		}
		if checkbigchar == true {
			if checklittlechar == true {
				if checknumber == true {
					if checksanyorat == true {
						fmt.Println("Valid")
					} else {
						fmt.Println("Invalid")
					}
				} else {
					fmt.Println("Invalid")
				}
			} else {
				fmt.Println("Invalid")
			}
		} else {
			fmt.Println("Invalid")
		}
	} else {
		fmt.Println("Invalid")
	}
}
