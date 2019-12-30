package commands

import "github.com/chyupa/fp700"

func FabricationNumber() {
	// print fabrication number; SAM Module
	fp700.SendCommand(71, "5\t")
}
