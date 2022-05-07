package problems

func NthMaximum(arr []int, n int) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	} else if len(arr) == 1 {
		return arr[0], true
	} else {
		var max int
		a := RemoveDuplicateIntegersInArray(arr)
		for i := 0; i < n; i++ {
			max, a = findMaxWithPop(a)
		}
		return max, true
	}
}

func findMaxWithPop(arr []int) (int, []int) {
	var max int
	var index int
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
			index = i
		}
	}
	if index == 0 {
		return max, arr[1:]
	} else if index == len(arr)-1 {
		return max, arr[0 : len(arr)-1]
	} else {
		arr1 := arr[0:index]
		arr1 = append(arr1, arr[index+1:]...)
		return max, arr1
	}
}
