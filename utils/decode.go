package utils

import (
	"errors"
	"fmt"
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
			return nil, errors.New(fmt.Sprintf("%d", errorCode))
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
