package commands

import (
	"fmt"
	"github.com/chyupa/fp700"
)

type PrintMfRequest struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	ReportType string `json:"reportType"`
	ByDate     bool   `json:"byDate"`
}

func PrintMf(mfReq PrintMfRequest) error {

	var command = 95
	payload := fmt.Sprintf("%s\t%s\t%s\t", mfReq.ReportType, mfReq.Start, mfReq.End)
	if mfReq.ByDate {
		command = 94
	}

	// initialize reading from EJ
	_, err := fp700.SendCommand(command, payload)
	if err != nil {
		return err
	}

	return nil
}
