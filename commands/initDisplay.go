package commands

import (
	"github.com/chyupa/fp700"
)

type InitDisplayRequest struct {
	FirstLine  string `json:"firstLine"'`
	SecondLine string `json:"secondLine"'`
}

func InitDisplay(idRequest InitDisplayRequest) {

	// reset display
	fp700.SendCommand(33, "")

	// set first line
	fp700.SendCommand(47, idRequest.FirstLine+"\t")

	// set second line
	fp700.SendCommand(35, idRequest.SecondLine+"\t")
}
