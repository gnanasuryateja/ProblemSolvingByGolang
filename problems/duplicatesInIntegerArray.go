package problems

func DuplicatesInIntegerArray(arr []int) int {
	countMap := make(map[int]int)
	for i := range arr {
		_, present := countMap[arr[i]]
		if present {
			return 1
		} else {
			countMap[arr[i]] = 1
		}
	}
	return -1
}
