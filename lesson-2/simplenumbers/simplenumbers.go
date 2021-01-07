package simplenumbers

import (
	"fmt"
	"math"
)

func simpleNumbers(N int64) ([]int64, error) {
	var i, j int64
	if N < 0 {
		return nil, fmt.Errorf("negative number")
	}
	numbers := make([]int64, 0, N)
	if N <= 1 {
		return numbers, nil
	}
	numbers = append(numbers, 2)
	for i = 3; i <= N; i += 2 {
		var sqrt = math.Sqrt(float64(i))
		if math.Ceil(sqrt) == sqrt {
			continue
		}

		simple := true
		for j = 2; j <= int64(sqrt); j++ {
			if i%j == 0 {
				simple = false
				break
			}
		}

		if simple == true {
			numbers = append(numbers, i)
		}
	}
	return numbers, nil
}
