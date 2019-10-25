package main

import "fmt"

func main() {
	var str = "Hello world"
	fmt.Println(str)

	str = "Testing"
	fmt.Println(str)

	//cannot use 1 (type int) as type string in assignment
	// str = 1
	// fmt.Println(str)

	var num int
	num = 2
	fmt.Println(num * num)

	num1 := 45
	num2 := 9
	fmt.Println(num1 / num2)

	var num3, num4 int = 3, 4
	fmt.Println(num3, num4)

	//default value:0
	var num5 int
	fmt.Println(num5)

	//default value:false
	var none bool
	fmt.Println(none)

}
