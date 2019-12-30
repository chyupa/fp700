package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
	"strings"
)

type CommandRequest struct {
	Code int `json:"code"`
	Payload string `json:"payload"`
}

type TransactionRequest struct {
	Commands []CommandRequest
}

type TransactionResponse struct {
	ErrorCode int
	IsOpen int
	Number int
	FNumberRep int
	FNumber int
	Items int
	Amount int
	Payed int
}
func Transaction(transactionRequest TransactionRequest) TransactionResponse {
	// check if receipt is open and cancel it
	LastReceipt(true)

	var transactionResponse TransactionResponse
	for _, command := range transactionRequest.Commands {
		//time.Sleep(time.Millisecond * 200)
		response, _ := fp700.SendCommand(command.Code, command.Payload)
		log.Println(response)
		if len(response) > 2 {
			msg := utils.DecodeMessage(response)
			log.Println(command, msg)
		}
	}

	checkLastReceipt, _ := fp700.SendCommand(76, "")
	if len(checkLastReceipt) > 2 {
		msg := utils.DecodeMessage(checkLastReceipt)
		split := strings.Split(msg, "\t")

		transactionResponse.ErrorCode, _ = strconv.Atoi(split[0])
		transactionResponse.IsOpen, _ = strconv.Atoi(split[1])
		transactionResponse.Number, _ = strconv.Atoi(split[2])
		transactionResponse.FNumberRep, _ = strconv.Atoi(split[3])
		transactionResponse.FNumber, _ = strconv.Atoi(split[4])
		transactionResponse.Items, _ = strconv.Atoi(split[5])
		transactionResponse.Amount, _ = strconv.Atoi(split[6])
		transactionResponse.Payed, _ = strconv.Atoi(split[7])
	}

	return transactionResponse
}
