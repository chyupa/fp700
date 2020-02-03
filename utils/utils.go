package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DecodedMessage struct {
	Raw    []byte
	Length int64
	Data   []byte
	Status []byte
}

func (dm *DecodedMessage) DecodeMessage(message []byte) ([]string, error) {
	startingIndex := IndexOf(message, 1)
	if startingIndex != -1 {
		message = append(message[startingIndex:])
	}

	if message[0] != 1 && message[len(message)-1] != 3 {
		return nil, errors.New("formatul mesajului este incorect")
	}
	dm.Raw = message
	length := reconstructHexFromByteArray(message[1:5])

	intLength, _ := strconv.ParseInt(length, 16, 0)
	dm.Length = intLength - 32

	dm.Data = message[10:dm.Length]

	statusSeparatorPosition := IndexOf(dm.Data, 0x04)
	if statusSeparatorPosition >= -1 {
		dm.Status = dm.Data[statusSeparatorPosition+1 : statusSeparatorPosition+8]
		msg := dm.Data[0:statusSeparatorPosition]

		if !dm.validateChecksum() {
			return nil, errors.New("nu am putut valida raspunsul")
		}

		split := strings.Split(string(msg), "\t")
		errorCode, _ := strconv.Atoi(split[0])
		if errorCode != -100003 && errorCode != -111015 && errorCode != -111016 && errorCode < 0 {
			return nil, errors.New(fmt.Sprintf("%d", int(math.Abs(float64(errorCode)))))
		}
		return split, nil
	}

	return nil, nil
}

func (dm *DecodedMessage) validateChecksum() bool {
	msgChecksum := dm.Raw[1 : len(dm.Raw)-5]
	var intChecksum []int
	for _, char := range msgChecksum {
		intChecksum = append(intChecksum, int(char))
	}

	sumChecksum := arraySum(intChecksum)
	hexChecksum := strconv.FormatInt(int64(sumChecksum), 16)

	expectedChecksum := hexToPrinterFormat(hexChecksum)

	givenChecksum := dm.Raw[dm.Length+1 : len(dm.Raw)-1]
	for idx, b := range givenChecksum {
		if int(b) != expectedChecksum[idx] {
			return false
		}
	}
	return true
}

func reconstructHexFromByteArray(byteArray []byte) string {
	hexString := ""

	for _, bit := range byteArray {
		b := bit - 0x30
		if b >= 0 {
			asd := strconv.FormatInt(int64(b), 16)
			hexString += asd
		}
	}

	return hexString
}

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

	bytes = append([]int{1}, bytes...)

	bytes = append(bytes, 3)

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
func ArrayContains(arr []byte, s int) bool {
	for _, a := range arr {
		if int(a) == s {
			return true
		}
	}
	return false
}

func calculateLength(payload string) []int {
	lengthBytes := 4
	seqBytes := 1
	cmdBytes := 4
	preambleBytes := 1
	offset := 32

	lengthDec := lengthBytes + seqBytes + cmdBytes + preambleBytes + len(payload) + offset
	lengthHex := strconv.FormatInt(int64(lengthDec), 16)
	return hexToPrinterFormat(lengthHex)
}

func IndexOf(byteArray []byte, bit byte) int {
	for index, b := range byteArray {
		if bit == b {
			return index
		}
	}

	return -1
}

var sequenceNumber = 32

func getSequenceNumber() int {
	if sequenceNumber > 254 {
		sequenceNumber = 32
	} else {
		sequenceNumber++
	}
	return sequenceNumber
}
