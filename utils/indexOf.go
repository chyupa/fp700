package utils

func IndexOf(byteArray []byte, bit byte) int {
	for index, b := range byteArray {
		if bit == b {
			return index
		}
	}

	return -1
}
