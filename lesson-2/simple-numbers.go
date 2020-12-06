package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var N, i, j int64
	fmt.Println("Введите число N")
	_, err := fmt.Scanln(&N)
	if err != nil {
		fmt.Println("Число введено не корректно")
		os.Exit(1)
	}

	if N < 0 {
		fmt.Println("Число должно быть положительным")
		os.Exit(1)
	}

	fmt.Println("-----------------------------")

	if N <= 1 {
		return
	}
	fmt.Println(2)
	for i = 3; i <= N; i = i + 2 {
		var sqrt = int64(math.Sqrt(float64(i)))
		if sqrt*sqrt == i {
			continue
		}
		simple := true
		for j = 2; j <= sqrt; j++ {
			if i%j == 0 {
				simple = false
				break
			}
		}

		if simple == true {
			fmt.Println(i)
		}
	}
}
