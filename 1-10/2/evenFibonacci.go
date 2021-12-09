package main

import "fmt"

func main() {
	var sum int
	fib := []int{1, 1} // the start of the fibonacci sequence, our first two numbers
	for {
		fibNext := fib[0] + fib[1] // add the last two indexes in the fibonacci sequence
		if fibNext >= 4000000 {
			break
		}
		if fibNext%2 == 0 { // check that it's divisible by 2 (even)
			sum += fibNext
		}
		fib[0] = fib[1] // move values to the left for next numbers in sequence
		fib[1] = fibNext
	}
	fmt.Println("The sum of even Fibonacci numbers below 4M is", sum)
}
