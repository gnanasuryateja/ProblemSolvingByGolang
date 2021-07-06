package problems

func MaxOccurancesInAString(data string) interface{} {
	if len(data) > 1 {
		countMap := make(map[string]int)
		for i := range data {
			_, present := countMap[string(data[i])]
			if present {
				countMap[string(data[i])] = countMap[string(data[i])] + 1
			} else {
				countMap[string(data[i])] = 1
			}
		}
		max := 1
		str := ""
		for key, val := range countMap {
			if max < val {
				str = key
			}
		}
		if str == "" {
			return nil
		}
		return str
	}
	return nil
}
