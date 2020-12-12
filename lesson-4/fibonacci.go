package main

import "fmt"

func main() {
	var n int64
	var memory = make(map[int64]int64)
	fmt.Println("Введите номер числа")
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	fmt.Println("Число фибоначчи: ", fib(n, memory))
}

func fib(n int64, memory map[int64]int64) int64 {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	val, ok := memory[n]
	if ok {
		fmt.Println("opt", val)
		return val
	}
	memory[n] = fib(n-2, memory) + fib(n-1, memory)
	if memory[n] < 0 {
		panic("int64 overflow")
	}
	return fib(n-2, memory) + fib(n-1, memory)
}
