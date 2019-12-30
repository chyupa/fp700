package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
)

type LastReceiptResponse struct {
	ErrorCode  int
	IsOpen     int
	Number     int
	FNumberRep int
	FNumber    int
	Items      int
	Amount     int
	Payed      int
}

func LastReceipt(shouldCancel bool) LastReceiptResponse {
	var decodedMessage = &utils.DecodedMessage{}
	var lastReceiptResponse LastReceiptResponse
	lastReceipt, _ := fp700.SendCommand(76, "")

	if len(lastReceipt) > 2 {
		msg, err := decodedMessage.DecodeMessage(lastReceipt)
		if err != nil {
			log.Println(err)
		}
		// receipt is open so we need to close it
		if msg[1] == "1" && shouldCancel {
			fp700.SendCommand(60, "")
		}
		lastReceiptResponse.ErrorCode, _ = strconv.Atoi(msg[0])
		lastReceiptResponse.IsOpen, _ = strconv.Atoi(msg[1])
		lastReceiptResponse.Number, _ = strconv.Atoi(msg[2])
		lastReceiptResponse.FNumberRep, _ = strconv.Atoi(msg[3])
		lastReceiptResponse.FNumber, _ = strconv.Atoi(msg[4])
		lastReceiptResponse.Items, _ = strconv.Atoi(msg[5])
		lastReceiptResponse.Amount, _ = strconv.Atoi(msg[6])
		lastReceiptResponse.Payed, _ = strconv.Atoi(msg[7])
	}

	return lastReceiptResponse
}
