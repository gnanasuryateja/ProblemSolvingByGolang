package problems

func PalindromeSubStrings(data string) map[string]int {
	var substrings []string
	for i := range data {
		for j := i; j < len(data); j++ {
			substrings = append(substrings, data[i:j+1])
		}
	}
	mp := make(map[string]int)
	for i := range substrings {
		if IsPalindrome(substrings[i]) && substrings[i] != "" {
			mp[substrings[i]] = 1
		}
	}
	return mp
}

func IsPalindrome(data string) bool {
	for i := 0; i < len(data)/2; i++ {
		if data[i] != data[len(data)-i-1] {
			return false
		}
	}
	return true
}
