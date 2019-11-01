package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	b := a
	fmt.Println(b)

	c := [5]int{1}
	fmt.Println(c)

	var two [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two[i][j] = i + j
		}
	}
	fmt.Println("two:", two)

	var strArray [5]string
	strArray[3] = "test"
	fmt.Println(strArray)

	strArray[2] = "lala"
	fmt.Println(strArray)

}
