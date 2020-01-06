package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
)

func RemainingZReports() int {
	var decodedMessage = &utils.DecodedMessage{}
	zReports, _ := fp700.SendCommand(68, "")
	zReportsDecoded, err := decodedMessage.DecodeMessage(zReports)
	if err != nil {
		log.Println(err)
	}

	remainingReports, _ := strconv.Atoi(zReportsDecoded[1])

	return remainingReports
}
