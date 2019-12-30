package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
)

type ServiceAmountRequest struct {
	amount string `json:"amount"`
}

func ServiceAmount(data ServiceAmountRequest) int {
	cmdResponse, _ := fp700.SendCommand(70, "0\t" + data.amount + "\t")

	msg := utils.DecodeMessage(cmdResponse)
	split := strings.Split(msg, "\t")

	errorCode, _ := strconv.Atoi(split[0])

	return errorCode
}
