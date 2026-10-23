package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

type Vertex struct {
	X int
	Y int
}

// Структура — это набор полей.

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

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	}
// 	return lim
// }

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: %v\n", i+1, z)
	}

	return z
}

// var c, python, java bool
// var i, j int = 1, 2

func main() {
	defer fmt.Println("world")

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

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("Final result:", Sqrt(2))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("Whe's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	i, j := 42, 2701

	p := &i //  указатель на i
	fmt.Println(p)
	fmt.Println(*p) // считываем i через указатель
	*p = 21         // устанавливаем i через указатель
	fmt.Println(i)  //  видим новое значение i

	p = &j         // указатель на j
	*p = *p / 37   // делим j через указатель
	fmt.Println(j) //  видим новое значение j

	fmt.Println(Vertex{1, 2})
}
