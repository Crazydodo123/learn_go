package main

import "fmt"

func fibonacci() func() int {
	fib_nums := []int{0, 1}
	return func() int {
		fib_nums = append(fib_nums, fib_nums[len(fib_nums)-1]+fib_nums[len(fib_nums)-2])
		return fib_nums[len(fib_nums)-3]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
