package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
)

func GetVat() string {
	var decodedMessage = &utils.DecodedMessage{}

	// ActiveVatGroups
	activeVatGroupsResponse, _ := fp700.SendCommand(50, "")
	activeVatGroupsDecoded, err := decodedMessage.DecodeMessage(activeVatGroupsResponse)
	if err != nil {
		log.Println(err)
	}
	return activeVatGroupsDecoded[2]
}

func GetVatChanges() int {
	var decodedMessage = &utils.DecodedMessage{}

	// nVatChanges
	nVatChangesResponse, _ := fp700.SendCommand(255, "nVatChanges\t\t\t")
	nVatChangesDecoded, err := decodedMessage.DecodeMessage(nVatChangesResponse)
	if err != nil {
		log.Println(err)
	}
	vatChanges, _ := strconv.Atoi(nVatChangesDecoded[1])
	return vatChanges
}

type SetVatRequest struct {
	Vat string `json:"vat"`
}

func SetVat(data SetVatRequest) (int, error) {
	var decodedMessage = &utils.DecodedMessage{}

	setVatResponse, _ := fp700.SendCommand(83, data.Vat+"\t")
	setVatDecoded, err := decodedMessage.DecodeMessage(setVatResponse)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	remainingChanges, err := strconv.Atoi(setVatDecoded[1])
	return remainingChanges, nil
}
