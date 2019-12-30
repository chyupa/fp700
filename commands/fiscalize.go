package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
	"strings"
)

type FiscalizeRequest struct {
	date string `json:"dateTime"`
	fabrication string `json:"fabricationNumber"`
	vat string `json:"vatValue"`
	tax string `json:"taxNo"`
	cif string `json:"cif"`
}

type FiscalizeResponse struct {
	ErrorCode int
}

func Fiscalize(data FiscalizeRequest) FiscalizeResponse {
	var fResponse FiscalizeResponse
	// set time
	timeResponse, _ := fp700.SendCommand(61, data.date + "\t")
	log.Println(timeResponse)

	// set serial number
	serialNumberResponse, _ := fp700.SendCommand(91, data.fabrication + "\t")
	log.Println(serialNumberResponse)

	// set VAT number
	vatResponse, _ := fp700.SendCommand(83, data.vat + "\t")
	log.Println(vatResponse)

	// set Fiscalization
	fiscalizeResponse, _ := fp700.SendCommand(72, data.tax + "\t" + data.cif + "\t")
	log.Println(fiscalizeResponse)
	msg := utils.DecodeMessage(fiscalizeResponse)
	split := strings.Split(msg, "\t")
	fResponse.ErrorCode, _ = strconv.Atoi(split[0])

	// disable close open receipt printer option
	fp700.SendCommand(255, "DsblKeyCloseReceipt\t\t1\t")
	// disable cancel receipt printer option
	fp700.SendCommand(255, "DsblKeyCancelReceipt\t\t1\t")
	// disable generation of fiscal memory reports
	fp700.SendCommand(255, "DsblKeyFmReports\t\t1\t")

	return fResponse
}
