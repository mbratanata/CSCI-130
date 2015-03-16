// hw5 project main.go
package main

import (
	"fmt"
	"import/math"
)

type likeClass struct { //struct type
	name  string
	class string
}

type stringMethod string

func (lc likeClass) teacher() {
	fmt.Println("THIS IS A METHOD ATTACHED TO A STRUCT: " + lc.name + " teaches " + lc.class)
}

func (m stringMethod) tryThisStringMethod() {
	fmt.Println("METHOD: THIS IS A METHOD ATTACHED TO A STRING")
}

func (t likeClass) teacherWithReturn() string {
	return "THIS IS A METHOD WITH RETURN: " + t.name + " teaches " + t.class
}

func main() {
	var class []string
	class = make([]string, 3, 5)
	class[0] = "CSCI 115"
	class[1] = "CSCI 130"
	class[2] = "MATH 77"
	class = append(class, "CSCI 112") //append to a slice
	fmt.Println(class[3])

	myTeacher := []string{"Todd Wilson", "Todd McLeod", "Adnan Sabuwala", "Jin Park"} //another way to append a slice
	myOtherTeacher := []string{"Ming Li", "Doreen DeLeon"}
	myTeacher = append(myTeacher, myOtherTeacher...)

	myTeacher = append(myTeacher[:5], myTeacher[6:]...) //delete from a slice
	fmt.Println()
	for _, entry := range myTeacher {
		fmt.Println(entry)
	}

	numberSlice := []int{5, 10, 15} //packages
	sumOfNumbers := math.Sum(numberSlice)
	fmt.Println(sumOfNumbers)

	fmt.Println()

	var message = likeClass{"Jin Park", "CSCI 112"} //methods attached to a struct
	message.teacher()

	fmt.Println()

	var input stringMethod = "Hi"
	input.tryThisStringMethod()

	fmt.Println()

	fmt.Println(message.teacherWithReturn())
}
