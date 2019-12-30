package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"

	"log"
	"strconv"
	"strings"
)

type ReportXResponse struct {
	ErrorCode int
	nRep int
	TotX int
	TotEXEPTAT int
	TotSInv int
	VatSInv int
}

func ReportX() ReportXResponse {
	response := ReportXResponse{}

	reply, err := fp700.SendCommand(69, "X\t")

	if err != nil {
		log.Fatal(err)
	}

	if len(reply) > 2 {
		msg := utils.DecodeMessage(reply)
		split := strings.Split(msg, "\t")

		response.ErrorCode, _ = strconv.Atoi(split[0])
		response.nRep, _ = strconv.Atoi(split[1])
		response.TotX, _ = strconv.Atoi(split[2])
		response.TotEXEPTAT, _ = strconv.Atoi(split[3])
		response.TotSInv, _ = strconv.Atoi(split[4])
		response.VatSInv, _ = strconv.Atoi(split[5])
	}

	return response
}
