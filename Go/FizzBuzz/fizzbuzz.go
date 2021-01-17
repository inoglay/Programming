package main

import "fmt"

func main() {
	var fizz, buzz, fizzbuzz, maxCount int
	fizz, buzz, fizzbuzz, maxCount = 3, 5, 15, 101

	i := 0
	for i < maxCount {
		if i%fizzbuzz == 0 {
			fmt.Println("fizzbuzz")
		} else if i%fizz == 0 {
			fmt.Println("fizz")
		} else if i%buzz == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

		i++
	}
}
