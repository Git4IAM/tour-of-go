package main

import "fmt"

var num, num2, num3 int = 10, 20, 30
var name string = "Jane Doe"
var pi float64 = 3.14
var numflaot float64 = 1.52

func main() {
	myvar1 := 39
	myvar2 := "GeeksforGeeks"
	myvar3 := 34.67

	fmt.Printf("myvar1: %d (%T)\n", myvar1, myvar1)
	fmt.Printf("myvar2: %s (%T)\n", myvar2, myvar2)
	fmt.Printf("myvar3: %f (%T)\n", myvar3, myvar3)
	/*myvar := 200
	fmt.Printf("myvar: %d (%T)\n", myvar)
	fmt.Println(num, num2)
	fmt.Println("name")
	fmt.Println(pi)*/
}
