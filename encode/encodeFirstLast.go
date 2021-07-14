package encode

func EncodeFirstLast(list []string) []string {
	result := make([]string, len(list))
	index := 0
	for _, value := range list {
		result[index] = value[:1]
		index ++
	}

	return result
}