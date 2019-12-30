package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strings"
)
type FootersDataResponse struct {
	Footer1 string
	Footer2 string
}

func FootersData() FootersDataResponse {
	var fdResponse FootersDataResponse

	// header
	footer1Response, _ := fp700.SendCommand(255, "Footer\t0\t\t")
	footer1Decoded := utils.DecodeMessage(footer1Response)
	footer1Split := strings.Split(footer1Decoded, "\t")
	fdResponse.Footer1 = footer1Split[1]

	footer2Response, _ := fp700.SendCommand(255, "Footer\t1\t\t")
	footer2Decoded := utils.DecodeMessage(footer2Response)
	footer2Split := strings.Split(footer2Decoded, "\t")
	fdResponse.Footer2 = footer2Split[1]

	return fdResponse
}

type SetFootersDataRequest struct {
	topLine string `json:"topLineValue"`
	bottomLine string `json:"bottomLineValue"`
}
func SetFootersData(data SetFootersDataRequest) {

	fp700.SendCommand(255, "Footer\t0\t" + data.topLine + "\t")
	fp700.SendCommand(255, "Footer\t1\t" + data.bottomLine + "\t")

}