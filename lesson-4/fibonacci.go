package main

import (
	"fmt"
	"os"
)

func main() {
	var n int64
	var memory = make(map[int64]int64)
	fmt.Println("Введите номер числа")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Введены не корректные данные!")
		os.Exit(1)
	}
	fibonacci, err := fib(n, memory)
	if err != nil {
		fmt.Println("Ошибка!", err.Error())
		os.Exit(1)
	}
	fmt.Println("Число фибоначчи: ", fibonacci)
}

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
