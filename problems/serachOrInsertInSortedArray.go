package problems

func SearchOrInsertInSortedArray(arr []int, key int) (string, int) {
	index := 0
	for i := range arr {
		if arr[i] == key {
			return "found in array, at position: ", i
		} else if arr[i] < key {
			index = index + 1
		} else if arr[i] > key {
			break
		}
	}
	return "to be inserted at position: ", index
}
