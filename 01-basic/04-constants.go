package main

import (
	"fmt"
	"math"
)

const name string = "yaohuiye"
const age int = 1
const gender = 1

func main() {
	fmt.Println(name, age, gender)

	const n = 5000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
