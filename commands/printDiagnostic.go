package commands

import "github.com/chyupa/fp700"

func PrintDiagnostic() {
	fp700.SendCommand(71, "0\t")
}
