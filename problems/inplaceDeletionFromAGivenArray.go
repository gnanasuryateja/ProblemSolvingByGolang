package problems

func InPlaceDeletion(arr []int, num int) []int {
	for i := range arr {
		if arr[i] == num && i == 0 {
			return arr[1:]
		} else if arr[i] == num && i == len(arr)-1 {
			return arr[:len(arr)-1]
		}
		if arr[i] == num {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return []int{}
}
