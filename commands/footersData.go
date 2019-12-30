package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
)

type FootersDataResponse struct {
	Footer1 string
	Footer2 string
}

func FootersData() FootersDataResponse {
	var decodedMessage = &utils.DecodedMessage{}
	var fdResponse FootersDataResponse

	// header
	footer1Response, _ := fp700.SendCommand(255, "Footer\t0\t\t")
	footer1Decoded, err := decodedMessage.DecodeMessage(footer1Response)
	if err != nil {
		log.Println(err)
	}
	fdResponse.Footer1 = footer1Decoded[1]

	footer2Response, _ := fp700.SendCommand(255, "Footer\t1\t\t")
	footer2Decoded, err := decodedMessage.DecodeMessage(footer2Response)
	if err != nil {
		log.Println(err)
	}
	fdResponse.Footer2 = footer2Decoded[1]

	return fdResponse
}

type SetFootersDataRequest struct {
	TopLine    string `json:"topLineValue"`
	BottomLine string `json:"bottomLineValue"`
}

func SetFootersData(data SetFootersDataRequest) {

	fp700.SendCommand(255, "Footer\t0\t"+data.TopLine+"\t")
	fp700.SendCommand(255, "Footer\t1\t"+data.BottomLine+"\t")

}
