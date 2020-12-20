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
	for i := 0; i < 8; i++ {
		if scanner.Scan() {
			line := scanner.Text()

			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}

			arr[i] = num

		} else {
			panic("you must input 8 numbers")
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
