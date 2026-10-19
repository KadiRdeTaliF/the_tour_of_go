package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

const Pi = 3.14

// var c, python, java bool
// var i, j int = 1, 2

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("the number is: ", rand.Intn(10))

	fmt.Println("now number is: ", math.Sqrt(10))

	fmt.Println(add(10, 14))

	// a, b := swap("hello", "world")
	// fmt.Println(a, b)

	fmt.Println(split(17))

	// var i int
	// fmt.Println(i, c, python, java)

	// var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)

	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"

	// fmt.Println(i, j, k, c, python, java)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
