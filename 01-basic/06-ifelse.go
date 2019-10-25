package main

import "fmt"

func main() {
	if 10/2 == 5 {
		fmt.Println("good result")
	} else {
		fmt.Println("bad result")
	}

	if num := 10; num < 11 {
		fmt.Println("< 11")
	} else if num == 11 {
		fmt.Println("= 11")
	} else {
		fmt.Println("> 11")
	}
}
