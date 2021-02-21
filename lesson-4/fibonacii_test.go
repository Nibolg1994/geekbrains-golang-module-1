package fibonacii

import (
	"fmt"
)

func ExampleFibonacci() {
	memory := make(map[int64]int64)
	fmt.Println(fib(8, memory))
	fmt.Println(fib(9, memory))
	fmt.Println(fib(10, memory))
	// Output: 13 <nil>
	// 21 <nil>
	// 34 <nil>
}
