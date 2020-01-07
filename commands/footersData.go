package commands

import (
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
)

type FootersDataResponse struct {
	Footer1 string
	Footer2 string
}

func FootersData() (FootersDataResponse, error) {
	var decodedMessage = &utils.DecodedMessage{}
	var fdResponse FootersDataResponse

	// header
	footer1Response, _ := fp700.SendCommand(255, "Footer\t0\t\t")
	footer1Decoded, err := decodedMessage.DecodeMessage(footer1Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return fdResponse, err
	}
	fdResponse.Footer1 = footer1Decoded[1]

	footer2Response, _ := fp700.SendCommand(255, "Footer\t1\t\t")
	footer2Decoded, err := decodedMessage.DecodeMessage(footer2Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return fdResponse, err
	}
	fdResponse.Footer2 = footer2Decoded[1]

	return fdResponse, nil
}

type SetFootersDataRequest struct {
	TopLine    string `json:"topLineValue"`
	BottomLine string `json:"bottomLineValue"`
}

func SetFootersData(data SetFootersDataRequest) error {
	var decodedMessage = &utils.DecodedMessage{}
	topFooter, _ := fp700.SendCommand(255, "Footer\t0\t"+data.TopLine+"\t")
	_, err := decodedMessage.DecodeMessage(topFooter)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return err
	}
	bottomFooter, _ := fp700.SendCommand(255, "Footer\t1\t"+data.BottomLine+"\t")
	_, err = decodedMessage.DecodeMessage(bottomFooter)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return err
	}

	return nil
}
