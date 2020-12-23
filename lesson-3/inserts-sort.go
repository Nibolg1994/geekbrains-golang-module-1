package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var arr [8]int64
	for i := 0; i < len(arr); i++ {
		if scanner.Scan() {
			line := scanner.Text()

			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				fmt.Println("Введены не корректные данные")
				os.Exit(1)
			}
			arr[i] = num
		} else {
			fmt.Println("Ошибка! Вы должны ввести 8 чисел")
			os.Exit(1)
		}
	}
	var slice = arr[:]
	fmt.Println("sort arr:", insertsSort(slice))
	fmt.Println("arr:", slice)
}

func insertsSort(slice []int64) []int64 {
	sliceCopy := make([]int64, len(slice))
	copy(sliceCopy, slice)
	for j := 1; j < len(sliceCopy); j++ {
		for i := j; i > 0 && sliceCopy[i-1] > sliceCopy[i]; i-- {
			sliceCopy[i], sliceCopy[i-1] = sliceCopy[i-1], sliceCopy[i]
		}

	}
	return sliceCopy
}
