package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("The sum of 3's and 5's below 1000 is", sum)
}
