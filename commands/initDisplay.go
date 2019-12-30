package commands

import (
	"github.com/chyupa/fp700"
)

func InitDisplay(firstLine string, secondLine string) {

	// reset display
	fp700.SendCommand(33, "")

	// set first line
	fp700.SendCommand(47, firstLine + "\t")

	// set second line
	fp700.SendCommand(35, secondLine + "\t")
}