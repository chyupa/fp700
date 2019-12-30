package utils

var sequenceNumber = 32

func getSequenceNumber() int {
	if sequenceNumber > 254 {
		sequenceNumber = 32
	} else {
		sequenceNumber++
	}
	return sequenceNumber
}
