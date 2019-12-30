package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
)

type CloseDmjeRequest struct {
	servicePass string
}

func CloseDmje(data CloseDmjeRequest) (int, error) {
	fp700.SendCommand(253, "0\t"+data.servicePass+"\t")
	closeResponse, _ := fp700.SendCommand(253, "2\t"+data.servicePass+"\t\t")
	var decodedMessage = &utils.DecodedMessage{}
	msg, err := decodedMessage.DecodeMessage(closeResponse)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	errorCode, _ := strconv.Atoi(msg[0])
	return errorCode, nil
}
