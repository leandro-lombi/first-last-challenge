package encode

func EncodeFirstLast(list []string) []string {
	first := []string{}
	last := []string{}
	result := []string{}
	count := 1
	for key, value := range list {
		first = append(first, value[:1])
		last = append(last, value[len(value)-1:])
		if count%2 == 0 && count < len(list) {
			result = append(result, first[key-1]+last[key])
		} else if count == len(list) {
			result = append(result, first[key-1]+last[key])
			result = append(result, first[key]+last[0])
		} else if count > 1 {
			result = append(result, first[key-1]+last[key])
		}
		count++
	}

	return result
}
