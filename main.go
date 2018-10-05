package main

import (
	"fmt"
)

func main() {
	numbers := (func() []int {
		temp := []int{}
		for i := 0; i <= 10; i++ {
			temp = append(temp, i)
		}
		return temp
	})()

	for _, n := range numbers {
		var val string
		if n%2 == 0 {
			val = "even"
		} else {
			val = "odd"
		}
		fmt.Println(n, "is "+val)
	}
}
