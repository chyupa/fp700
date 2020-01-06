package commands

import (
	"errors"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
)

func Time() (string, error) {
	var decodedMessage = &utils.DecodedMessage{}
	reply, err := fp700.SendCommand(62, "")
	if err != nil {
		log.Println(err)
		return "", err
	}

	if len(reply) > 2 {
		msg, err := decodedMessage.DecodeMessage(reply)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return msg[1], nil
	}

	return "", errors.New("something went wrong")
}

type SetTimeRequest struct {
	Time string `json:"time"`
}

func SetTime(data SetTimeRequest) string {
	var decodedMessage = &utils.DecodedMessage{}

	setTimeResponse, _ := fp700.SendCommand(61, data.Time+"\t")
	setTimeDecoded, err := decodedMessage.DecodeMessage(setTimeResponse)
	if err != nil {
		log.Println(err)
	}

	return setTimeDecoded[0]
}
