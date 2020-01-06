package commands

import (
	"fmt"
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

type SetOperatorNameRequest struct {
	Name string `json:"name"`
}

func SetOperatorName(data SetOperatorNameRequest) (string, error) {
	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, fmt.Sprintf("OperName\t0\t%s\t", data.Name))

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[0], nil
}

func GetOperatorPassword() (string, error) {

	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, "OperPasw\t0\t\t")

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[1], nil
}

type SetOperatorPasswordRequest struct {
	OldPass string `json:"oldPass"`
	NewPass string `json:"newPass"`
}

func SetOperatorPassword(data SetOperatorPasswordRequest) error {
	var decodedMessage utils.DecodedMessage
	payload := fmt.Sprintf("1\t%s\t%s\t", data.OldPass, data.NewPass)
	cmd, _ := fp700.SendCommand(101, payload)

	_, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return err
	}

	return nil
}
