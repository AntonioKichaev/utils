package utils

func Contains(key string, sl []string) bool {
	for _, el := range sl {
		if el == key {
			return true
		}
	}
	return false
}
