package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
)

type CloseDmjeRequest struct{
	servicePass string
}

func CloseDmje(data CloseDmjeRequest) int {
	fp700.SendCommand(253, "0\t" + data.servicePass + "\t")
	closeResponse, _ := fp700.SendCommand(253, "2\t" + data.servicePass + "\t\t")

	msg := utils.DecodeMessage(closeResponse)
	split := strings.Split(msg, "\t")

	errorCode, _ := strconv.Atoi(split[0])
	return errorCode
}
