package fibonacii

import (
	"fmt"
)

func fib(n int64, memory map[int64]int64) (int64, error) {
	if n == 1 {
		return 0, nil
	}
	if n == 2 {
		return 1, nil
	}
	val, ok := memory[n]
	if ok {
		return val, nil
	}

	f1, err := fib(n-2, memory)

	if err != nil {
		return 0, err
	}

	f2, err := fib(n-1, memory)
	if err != nil {
		return 0, err
	}

	memory[n] = f1 + f2
	if memory[n] < 0 {
		return 0, fmt.Errorf("int64 overflow")
	}

	return memory[n], nil
}
