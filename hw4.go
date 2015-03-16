// hw4 project main.go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ { //loop
		for i > 6 {
			break
		}
		fmt.Println(i)
	}

	x := 0
	for x < 10 {
		if x >= 7 {
			break //break
		}
		if x%2 == 0 {
			x++
			continue //continue
		}
		fmt.Println(x)
		x++
	}

	thisSlice := []int{1, 5, 15, 20, 25, 30} //looping over a range
	for _, entry := range thisSlice {
		fmt.Println(entry)
	}

	myClass := map[string]string{ //maps
		"CSCI 115": "Todd Wilson",
		"CSCI 130": "Todd McLeod",
		"MATH 77":  "Adnan Sabuwala",
		"CSCI 112": "Jin Park",
	}
	fmt.Println(myClass["CSCI 130"])

	for i, entry := range thisSlice { //slices
		fmt.Println(i, " - ", entry)
	}
}
