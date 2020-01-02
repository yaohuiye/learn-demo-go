package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusMore(a, b, c int) int {
	return a + b + c
}

func test1() (int, int) {
	return 3, 8
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	result := plus(2, 4)
	fmt.Println(result)

	result = plusMore(2, 4, 6)
	fmt.Println(result)

	a, b := test1()
	fmt.Printf("%d,%d\n", a, b)

	_, c := test1()
	fmt.Println(c)

	sum(1, 3, 4)
	nums1 := []int{1, 2, 3, 4, 5, 6}
	sum(nums1...)
}
