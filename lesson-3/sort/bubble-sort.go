package sort

func bubbleSort(slice []int64) []int64 {
	sliceCopy := make([]int64, len(slice))
	copy(sliceCopy, slice)
	for i := 0; i < len(sliceCopy); i++ {
		for j := 0; j < len(sliceCopy)-i-1; j++ {
			if sliceCopy[j] > sliceCopy[j+1] {
				sliceCopy[j], sliceCopy[j+1] = sliceCopy[j+1], sliceCopy[j]
			}
		}
	}
	return sliceCopy
}
