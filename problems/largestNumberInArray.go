package problems

func LargestInArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
			}
		}
		return max
	}
}
