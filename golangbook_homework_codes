// golangbookhw project main.go
package main

import (
	"fmt"
	"math"
	"time"
)

// Chapter 4
func fahrenheitToCelsius() {
	fmt.Print("Enter a Fahrenheit temperature: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * (5.0 / 9.0)
	fmt.Println(output)
}

func feetToMeter() {
	fmt.Print("Enter a number in feet: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println(output)
}

// Chapter 5
func div3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Chapter 6
func smallest() int {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	min := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

// Chapter 7
func half(x int) (int, bool) {
	return (x / 2), (x%2 == 0)
}

func greatest(x ...int) int {
	greatest := x[0]
	for _, v := range x {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// Chapter 8
func swap(x *int, y *int) {
	temp := *y
	*y = *x
	*x = temp
}

// Chapter 9
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	perimeter() float64
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

func (r *Rectangle) perimeter() float64 {
	return (2 * ((r.y2 - r.y1) + (r.x2 - r.x1)))
}

// Chapter 10
func sleepFunc(n int) int {
	<-time.After(time.Second)
	return n
}

// Chapter 11
func Min(x []float64) float64 {
	min := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}
func Max(x []float64) float64 {
	max := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func main() {
	// Chapter 4
	// fahrenheitToCelsius()
	// feetToMeter()

	// Chapter 5
	// div3()
	// fizzBuzz ()

	// Chapter 6
	// fmt.Println(smallest())

	// Chapter 7
	// fmt.Println(half(1))
	// fmt.Println(half(2))
	// fmt.Println(greatest(1, 2, 3, 4, 5))
	/*nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())*/
	// fmt.Println(fib(0))
	/*x := 1
	y := 2

	// Chapter 8
	swap(&x, &y)*/
	// fmt.Println(x, y)

	// Chapter 9
	/*crc := Circle{0, 0, 5}
	fmt.Println("Perimeter of the circle is ", crc.perimeter())
	rec := Rectangle{0, 0, 4, 5}
	fmt.Println("Perimeter of the rectangle is ", rec.perimeter())*/

	// Chapter 10
	// fmt.Println("Your sleep time in seconds is: ", sleepFunc(100), "seconds.")
	// fmt.Println("Your sleep time in seconds is: ", sleepFunc(100), "seconds.")

	// Chapter 11
	/*x := []float64{0.5, 24.5, 3.2, 8.95, 2.45}
	fmt.Println(Min(x))
	fmt.Println(Max(x))*/
}
