package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
)

type CancelReceiptResponse struct {
	ErrorCode int
}

func CancelReceipt() (CancelReceiptResponse, error) {
	response, _ := fp700.SendCommand(60, "")
	var decodedMessage = &utils.DecodedMessage{}

	if len(response) > 1 {
		msg, err := decodedMessage.DecodeMessage(response)
		if err != nil {
			log.Println(err)
			return CancelReceiptResponse{}, err
		}
		if len(msg) > 0 {
			errorCode, _ := strconv.Atoi(msg[0])
			return CancelReceiptResponse{
				ErrorCode: errorCode,
			}, nil
		}
	}

	return CancelReceiptResponse{
		ErrorCode: 1,
	}, nil

}
