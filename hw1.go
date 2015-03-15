// hw1 project main.go
package main

import "fmt"

type hello string //type

type likeClass struct { //struct
	language string
	compiler string
}

func main() {
	var math hello = "Mathematics"
	fmt.Println(math)

	var input string
	input = "This is my first Go program."

	var inputAddress *string = &input //pointers
	var sameInput *string = &input

	fmt.Println(input)
	fmt.Println(inputAddress)
	fmt.Println(*sameInput)

	*sameInput = "This language is pretty cool."

	fmt.Println(*sameInput)

	//-------------------//

	var intro = likeClass{compiler: "LiteIDE", language: "GO"}

	fmt.Println(intro.language)
	fmt.Println(intro.compiler) // you can't combine the intro.language and intro.compiler in one single fmt.Println
}
