package commands

import (
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
)

type CommandRequest struct {
	Code    int    `json:"code"`
	Payload string `json:"payload"`
}

type TransactionRequest struct {
	Commands []CommandRequest
}

type TransactionResponse struct {
	ErrorCode  int
	IsOpen     int
	Number     int
	FNumberRep int
	FNumber    int
	Items      int
	Amount     int
	Payed      int
}

func Transaction(transactionRequest TransactionRequest) (TransactionResponse, error) {
	var decodedMessage = &utils.DecodedMessage{}
	// check if receipt is open and cancel it
	LastReceipt(true)

	var transactionResponse TransactionResponse
	for _, command := range transactionRequest.Commands {
		response, _ := fp700.SendCommand(command.Code, command.Payload)
		if len(response) > 2 {
			_, err := decodedMessage.DecodeMessage(response)
			if err != nil {
				fmt.Println(err)
				logger.Error.Println(err)
				return transactionResponse, err
			}
		}
	}

	checkLastReceipt, _ := fp700.SendCommand(76, "")
	if len(checkLastReceipt) > 2 {
		msg, err := decodedMessage.DecodeMessage(checkLastReceipt)
		if err != nil {
			fmt.Println(err)
			logger.Error.Println(err)
			return transactionResponse, err
		}

		transactionResponse.ErrorCode, _ = strconv.Atoi(msg[0])
		transactionResponse.IsOpen, _ = strconv.Atoi(msg[1])
		transactionResponse.Number, _ = strconv.Atoi(msg[2])
		transactionResponse.FNumberRep, _ = strconv.Atoi(msg[3])
		transactionResponse.FNumber, _ = strconv.Atoi(msg[4])
		transactionResponse.Items, _ = strconv.Atoi(msg[5])
		transactionResponse.Amount, _ = strconv.Atoi(msg[6])
		transactionResponse.Payed, _ = strconv.Atoi(msg[7])
	}

	return transactionResponse, nil
}
