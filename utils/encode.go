package utils

import (
	"strconv"
)

func EncodeMessage(command int, payload string) []byte {
	var bytes []int
	sequenceNumber = getSequenceNumber()

	if len(payload) < 1 {
		payload = ""
	}

	// LEN
	length := calculateLength(payload)

	bytes = append(bytes, length...)

	// SEQ
	bytes = append(bytes, sequenceNumber)

	// CMD
	commandHex := strconv.FormatInt(int64(command), 16)
	com := hexToPrinterFormat(string(commandHex))

	bytes = append(bytes, com...)

	// DATA
	for _, character := range payload {
		bytes = append(bytes, int(character))
	}

	bytes = append(bytes, 5)

	// BCC / Control sum
	sumChecksum := arraySum(bytes)
	hexChecksum := strconv.FormatInt(int64(sumChecksum), 16)

	checksum := hexToPrinterFormat(hexChecksum)

	bytes = append(bytes, checksum...)

	bytes = append(bytes, 3)

	bytes = append([]int{1}, bytes...)

	return convertToBytes(bytes)
}

func arraySum(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return sum
}

func convertToBytes(arr []int) []byte {
	var bytes []byte

	for _, val := range arr {

		bytes = append(bytes, byte(val))
	}

	return bytes
}
