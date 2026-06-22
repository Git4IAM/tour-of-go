package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

var num, num2, num3 int = 10, 20, 30
var name string = "Jane Doe"
var pi float64 = 3.14
var numflaot float64 = 1.52

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

type Vertex struct {
	X, Y int
}

var (
	v1       = Vertex{1, 2}
	v2       = Vertex{X: 1}
	v3       = Vertex{}
	literals = &Vertex{1, 2}
)

func main() {
	myvar1 := 39
	myvar2 := "GeeksforGeeks"
	myvar3 := 34.67

	fmt.Printf("myvar1: %d (%T)\n", myvar1, myvar1)
	fmt.Printf("myvar2: %s (%T)\n", myvar2, myvar2)
	fmt.Printf("myvar3: %f (%T)\n", myvar3, myvar3)
	myvar := 200
	fmt.Printf("myvar: %d (%T)\n", myvar)
	fmt.Println(num, num2)
	fmt.Println("name")
	fmt.Println(pi)

	const Kind = 48
	fmt.Printf("Type of Kind is: %T ", Kind)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 10),
	)

	x := 4.0
	fmt.Println("Yours: ", Sqrt(x))
	fmt.Println("math.sqrt: ", math.Sqrt(x))

	//switch
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}

	i, j := 42, 2701

	p := &i         //p point to i
	fmt.Println(*p) //read i through pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //pointer to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) //see new value of j

	v := Vertex{1, 2}
	v.X = 4 //get field of struct
	fmt.Println(v.X)
	fmt.Println(Vertex{1, 2})

	fmt.Println(v1, literals, v2, v3)

	var arr [10]string //array
	arr[0] = "sakurazaka"
	arr[1] = "hinatazaka"
	arr[2] = "nogizaka"

	fmt.Println(arr[2], arr[0])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var slice []int = primes[1 : 4+1] //slices [0, 1, 2, 3, 4, 5]
	fmt.Println(primes)
	fmt.Println(slice)
}
