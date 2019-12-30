package commands

import (
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
)

func GetOperatorPassword() (string, error) {

	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, "OperPasw\t0\t\t")

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[1], nil
}

func SetOperatorPassword(oldPass string, newPass string) error {
	var decodedMessage utils.DecodedMessage
	payload := fmt.Sprintf("1\t%s\t%s\t", oldPass, newPass)
	cmd, _ := fp700.SendCommand(101, payload)

	_, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return err
	}

	return nil
}
