package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700-server/utils/logger"
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

type CancelReceiptResponse struct {
	ErrorCode int
}

func CancelReceipt() (CancelReceiptResponse, error) {
	response, _ := fp700.SendCommand(60, "")
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

func LastReceipt(shouldCancel bool) (LastReceiptResponse, error) {
	var lastReceiptResponse LastReceiptResponse
	lastReceipt, _ := fp700.SendCommand(76, "")

	if len(lastReceipt) > 2 {
		msg, err := decodedMessage.DecodeMessage(lastReceipt)
		if err != nil {
			logger.Error.Println(err)
			return lastReceiptResponse, err
		}
		// receipt is open so we need to close it
		if msg[1] == "1" && shouldCancel {
			cancel, _ := fp700.SendCommand(60, "")
			_, err := decodedMessage.DecodeMessage(cancel)
			if err != nil {
				return lastReceiptResponse, err
			}
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

	return lastReceiptResponse, nil
}

type ServiceAmountRequest struct {
	amount string `json:"amount"`
}

func ServiceAmount(data ServiceAmountRequest) int {
	cmdResponse, _ := fp700.SendCommand(70, "0\t"+data.amount+"\t")

	msg, err := decodedMessage.DecodeMessage(cmdResponse)
	if err != nil {
		log.Println(err)
	}

	errorCode, _ := strconv.Atoi(msg[0])

	return errorCode
}
