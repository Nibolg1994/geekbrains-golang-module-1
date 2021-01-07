package sort

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
