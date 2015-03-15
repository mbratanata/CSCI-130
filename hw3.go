// hw3 project main.g
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

type TypeFunction func(string)

func Comments(comment LikeClass, tryThis TypeFunction) {
	MyClass, MyTime := NewFunction(comment.class, comment.time)
	tryThis(MyClass)
	tryThis(MyTime)
}

func NewFunction(class, time string) (message1 string, message2 string) { //multiple returns and function2
	message1 = class + " meets at " + time
	message2 = class + " is so interesting!"
	return
}

func myPrint(s string) {
	fmt.Print(s)
}

func myPrintln(s string) {
	fmt.Println(s)
}

func myPrintFunction(custom string) TypeFunction { // funcs that return funcs output
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func checkThis(check bool) { //if
	if check {
		fmt.Println("I have a class!")
	}
	fmt.Println("I don't have a class!")
}

func checkThis2(check2 bool) { //if_else
	if check2 {
		fmt.Println("I have a class!")
	} else {
		fmt.Println("I don't have a class!")
	}
}

func main() {
	var x = LikeClass{}
	x.class = "CSCI 130"
	x.time = "MW 1:00-02:50"

	fmt.Println(x.language)
	fmt.Println(x.compiler)

	Comments(x, myPrintln)

	a := LikeClass{}
	a.class = "MATH 77"
	a.time = "MTWTH 03:00-03:50"
	Comments(a, func(s string) { //funcs that accept funcs as arguments output
		fmt.Println(s)
	})

	b := LikeClass{}
	b.class = "CSCI 112"
	b.time = "TTH 11:00-12:15"
	Comments(b, myPrintFunction(" Good")) // funcs that return funcs output

	checkThis(false) //if output
	checkThis2(true) //if_else output

	fmt.Println()

	switch "CSCI 115" { //switch
	case "CSCI 112":
		fmt.Println("Jin Park")
	case "CSCI 115":
		fmt.Println("Todd WIlson")
	}
}
