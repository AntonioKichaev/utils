package utils

func InSlice(key string, sl []string) bool {
	for _, el := range sl {
		if el == key {
			return true
		}
	}
	return false
}

func IntInSlice3(key int, sl []int) bool {
	for _, el := range sl {
		if el == key {
			return true
		}
	}
	return false
}
