package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	_, prs1 := m["k1"]
	fmt.Println("prs:", prs1)

	n := map[string]int{"apple": 2, "banana": 3}
	fmt.Println("map:", n)
}
