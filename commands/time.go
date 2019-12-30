package commands

import (
	"errors"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
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
		if len(msg) > 1 {
			errorCode, _ := strconv.Atoi(msg[0])
			if errorCode < -1 {
				return "", errors.New(msg[0])
			}
			return msg[1], nil
		}
	}

	return "", errors.New("something went wrong")
}
