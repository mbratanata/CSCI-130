// hw2 project main.go
package main

import "fmt"

type LikeClass struct { //struct type
	language string
	compiler string
	input    string
	message  string
	class    string
	time     string
}

const ( //constants
	message1 = "Go is a cool language."
	message2 = "Go is not a proper language, which is false."
)

func Comments(comment LikeClass) { //function1
	class, _ := NewFunction(comment.class, comment.time) //unused returns without error
	fmt.Println(class)
	fmt.Println(comment.input)
	fmt.Println(comment.message)
	fmt.Println(NewFunction(comment.class, comment.time))
}

func NewFunction(class, time string) (string, int) { //multiple returns and function2
	count := 1
	return class + " meets at " + time, count
}

func main() {
	fmt.Println(message1)
	fmt.Println(message2)
	var x = LikeClass{}
	x.language = "Go"
	x.compiler = "LiteIDE"

	x.input = "This is my second Go program."
	x.message = "This language is pretty cool."

	x.class = "CSCI 130"
	x.time = "MW 1:00-02:50"

	fmt.Println(x.language)
	fmt.Println(x.compiler)

	Comments(x)

	a := LikeClass{}
	a.class = "MATH 77"
	a.time = "MTWTH 03:00-03:50"
	Comments(a)
}
g
