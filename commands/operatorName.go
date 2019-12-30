package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
)

func GetOperatorName() (string, error) {

	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, "OperName\t0\t\t")

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[1], nil
}
