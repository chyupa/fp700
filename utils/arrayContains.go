package utils

func ArrayContains(arr []byte, s int) bool {
	for _, a := range arr {
		if int(a) == s {
			return true
		}
	}
	return false
}
