package main

import "fmt"

func main() {
	var sum int
	fib := []int{1, 1}
	for {
		fibNext := fib[0] + fib[1]
		if fibNext >= 4000000 {
			break
		}
		if fibNext%2 == 0 {
			sum += fibNext
		}
		fib[0] = fib[1]
		fib[1] = fibNext
	}
	fmt.Println("The sum of even Fibonacci numbers below 4M is", sum)
}
