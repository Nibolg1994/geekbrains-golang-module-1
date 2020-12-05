package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b, res float64
	var inputa, inputb string
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&inputa)
	a, err := strconv.ParseFloat(inputa, 64)
	if err != nil {
		fmt.Println("Введены не корректные данные")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&inputb)
	b, err = strconv.ParseFloat(inputb, 64)
	if err != nil {
		fmt.Println("Введены не корректные данные")
		os.Exit(1)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "%":
		if a == float64(int64(a)) && b == float64(int64(b)) {
			res = float64(int64(a) % int64(b))
		} else {
			fmt.Println("В этой операции должны быть целые числа")
			os.Exit(1)
		}
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
