package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
)

type CancelReceiptResponse struct {
	ErrorCode int
}

func CancelReceipt() CancelReceiptResponse {
	response, _ := fp700.SendCommand(60, "")

	if len(response) > 1 {
		msg := utils.DecodeMessage(response)

		split := strings.Split(msg, "\t")
		if len(split) > 0 {
			errorCode, _ := strconv.Atoi(split[0])
			return CancelReceiptResponse{
				ErrorCode: errorCode,
			}
		}
	}

	return CancelReceiptResponse{
		ErrorCode: 1,
	}

}
