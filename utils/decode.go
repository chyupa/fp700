package utils

import (
	"strconv"
)

func DecodeMessage(message []byte) string {
	length := reconstructHexFromByteArray(message[1:5])

	intLength, _ := strconv.ParseInt(length, 16, 0)
	LEN := intLength - 32

	DATA := message[10:LEN]

	statusSeparatorPosition := IndexOf(DATA, 0x04)
	if statusSeparatorPosition >= -1 {
		msg := DATA[0:statusSeparatorPosition]

		return string(msg)
	}

	return ""
}

func reconstructHexFromByteArray(byteArray []byte) string {
	var hexString = ""

	for _, bit := range byteArray {
		b := bit - 0x30
		if b >= 0 {
			asd := strconv.FormatInt(int64(b), 16)
			hexString += asd
		}
	}

	return hexString
}

func IndexOf(byteArray []byte, bit byte) int {
	for index, b := range byteArray {
		if bit == b {
			return index
		}
	}

	return -1
}

func ArrayContains(arr []byte, s int) bool {
	for _, a := range arr {
		if int(a) == s {
			return true
		}
	}
	return false
}
