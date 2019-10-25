package main

import (
	"fmt"
	"time"
)

func main() {
	i := 12

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 12:
		fmt.Println("twelve")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("sleep more")
	default:
		fmt.Println("sleep less")
	}

	fmt.Println(time.Now())

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	default:
		fmt.Println("afternoon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("don't know")
		}
	}

	whatAmI(true)

	whatAmI("test")
}
