package utils

import (
	"strconv"
)

func hexToPrinterFormat(hexString string) []int {
	var lengthDecArray []int

	for _, c := range hexString {
		// fmt.Println("rangechar", c, string(c))
		hextByte := "3" + string(c)
		// fmt.Println("range hex byte", hextByte)
		converted, _ := strconv.ParseInt(hextByte, 16, 0)
		// fmt.Println("range converted", converted)
		lengthDecArray = append(lengthDecArray, int(converted))
	}

	remainingBits := 4 - len(hexString)
	hextByte, _ := strconv.ParseInt("30", 16, 0)

	for i := 0; i < remainingBits; i++ {
		// fmt.Println("remaining ", i, hextByte)
		lengthDecArray = append([]int{int(hextByte)}, lengthDecArray...)
	}

	return lengthDecArray
}
