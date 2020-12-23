package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, b, res float64
	var inputa, inputb string
	var op string
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&inputa)
	if err != nil {
		fmt.Println("Введены не корректные данные")
		os.Exit(1)
	}
	a, err = strconv.ParseFloat(inputa, 64)
	if err != nil {
		fmt.Println("Введены не корректные данные")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&inputb)
	if err != nil {
		fmt.Println("Введены не корректные данные")
		os.Exit(1)
	}
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
		if b == 0 {
			fmt.Println("В этой операции деление на ноль запрещено")
			os.Exit(1)
		}
		res = a / b
	case "%":
		if b == 0 {
			fmt.Println("В этой операции деление на ноль запрещено")
			os.Exit(1)
		}
		if a != math.Trunc(a) || b != math.Trunc(b) {
			fmt.Println("В этой операции должны быть целые числа")
			os.Exit(1)
		}
		res = float64(int64(a) % int64(b))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
