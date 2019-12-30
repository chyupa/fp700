package utils

import (
	"strconv"
)

var offset = 32

func calculateLength(payload string) []int {
	var lengthBytes = 4
	var seqBytes = 1
	var cmdBytes = 4
	var preambleBytes = 1

	lengthDec := lengthBytes + seqBytes + cmdBytes + preambleBytes + len(payload) + offset
	lengthHex := strconv.FormatInt(int64(lengthDec), 16)
	return hexToPrinterFormat(lengthHex)
}
