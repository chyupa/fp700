package commands

import (
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
)

type PrintEjRequest struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	ByDate bool   `json:"byDate"`
}

func PrintEj(pEjReq PrintEjRequest) error {

	var decodedMessage = &utils.DecodedMessage{}

	command := 125
	payload := fmt.Sprintf("13\t%s\t%s\t", pEjReq.Start, pEjReq.End)
	if pEjReq.ByDate {
		command = 124
		payload = fmt.Sprintf("%s\t%s\t2\t", pEjReq.Start, pEjReq.End)

		initialCommand, err := fp700.SendCommand(command, payload)
		if err != nil {
			log.Println(err)
		}

		initialCommandResponse, err := decodedMessage.DecodeMessage(initialCommand)
		if err != nil {
			return err
		}

		pEjReq.ByDate = false
		pEjReq.Start = initialCommandResponse[4]
		pEjReq.End = initialCommandResponse[6]

		return PrintEj(pEjReq)
	}

	_, err := fp700.SendCommand(command, payload)
	fp700.SendCommand(46, "")
	return err
}
