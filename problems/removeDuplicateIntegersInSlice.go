package problems

func RemoveDuplicateIntegersInArray(arr []int) []int {
	mp := make(map[int]int)
	var a []int
	for i := range arr {
		_, present := mp[arr[i]]
		if !present {
			a = append(a, arr[i])
			mp[arr[i]] = 1
		} else {
			mp[arr[i]] = mp[arr[i]] + 1
		}
	}
	return a
}
