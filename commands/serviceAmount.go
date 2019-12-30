package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
)

type ServiceAmountRequest struct {
	amount string `json:"amount"`
}

func ServiceAmount(data ServiceAmountRequest) int {
	var decodedMessage = &utils.DecodedMessage{}
	cmdResponse, _ := fp700.SendCommand(70, "0\t"+data.amount+"\t")

	msg, err := decodedMessage.DecodeMessage(cmdResponse)
	if err != nil {
		log.Println(err)
	}

	errorCode, _ := strconv.Atoi(msg[0])

	return errorCode
}
