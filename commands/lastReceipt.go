package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
)

type LastReceiptResponse struct {
	ErrorCode int
	IsOpen int
	Number int
	FNumberRep int
	FNumber int
	Items int
	Amount int
	Payed int
}

func LastReceipt(shouldCancel bool) LastReceiptResponse {
	var lastReceiptResponse LastReceiptResponse
	lastReceipt, _ := fp700.SendCommand(76, "")

	if len(lastReceipt) > 2 {
		msg := utils.DecodeMessage(lastReceipt)
		split := strings.Split(msg, "\t")
		// receipt is open so we need to close it
		if split[1] == "1" && shouldCancel {
			fp700.SendCommand(60, "")
		}
		lastReceiptResponse.ErrorCode, _ = strconv.Atoi(split[0])
		lastReceiptResponse.IsOpen, _ = strconv.Atoi(split[1])
		lastReceiptResponse.Number, _ = strconv.Atoi(split[2])
		lastReceiptResponse.FNumberRep, _ = strconv.Atoi(split[3])
		lastReceiptResponse.FNumber, _ = strconv.Atoi(split[4])
		lastReceiptResponse.Items, _ = strconv.Atoi(split[5])
		lastReceiptResponse.Amount, _ = strconv.Atoi(split[6])
		lastReceiptResponse.Payed, _ = strconv.Atoi(split[7])
	}

	return lastReceiptResponse
}
