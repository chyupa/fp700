package commands

import (
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"

	"log"
	"strconv"
)

type ReportXResponse struct {
	ErrorCode  int
	nRep       int
	TotX       int
	TotEXEPTAT int
	TotSInv    int
	VatSInv    int
}

func PrintReport(reportType string) ReportXResponse {
	var decodedMessage = &utils.DecodedMessage{}
	response := ReportXResponse{}

	payload := fmt.Sprintf("%s\t", reportType)
	reply, err := fp700.SendCommand(69, payload)

	if err != nil {
		log.Fatal(err)
	}

	if len(reply) > 2 {
		msg, err := decodedMessage.DecodeMessage(reply)
		if err != nil {
			log.Println(err)
		}

		response.ErrorCode, _ = strconv.Atoi(msg[0])
		response.nRep, _ = strconv.Atoi(msg[1])
		response.TotX, _ = strconv.Atoi(msg[2])
		response.TotEXEPTAT, _ = strconv.Atoi(msg[3])
		response.TotSInv, _ = strconv.Atoi(msg[4])
		response.VatSInv, _ = strconv.Atoi(msg[5])
	}

	return response
}
