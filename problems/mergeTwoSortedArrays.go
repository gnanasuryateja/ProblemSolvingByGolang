package problems

func MergeTwoSortedArrays(arr1 []int, arr2 []int) []int {
	var arr []int
	var cnt1, cnt2 int
	for cnt1 < len(arr1) && cnt2 < len(arr2) {
		if arr1[cnt1] < arr2[cnt2] {
			arr = append(arr, arr1[cnt1])
			cnt1 = cnt1 + 1
		} else {
			arr = append(arr, arr2[cnt2])
			cnt2 = cnt2 + 1
		}
	}
	if cnt1 < len(arr1) {
		arr = append(arr, arr1[cnt1:]...)
	} else if cnt2 < len(arr2) {
		arr = append(arr, arr2[cnt2:]...)
	}
	return arr
}
