package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
	"strings"
)

type TimeResponse struct {
	ErrorCode int
	Time string
}
func Time() *TimeResponse  {

	reply, err := fp700.SendCommand(62, "")
	if err != nil {
		log.Fatal(err)
	}

	if len(reply) > 2 {
		msg := utils.DecodeMessage(reply)
		split := strings.Split(msg, "\t")
		if len(split) > 1 {
			errorCode, _ := strconv.Atoi(split[0])
			return &TimeResponse{
				ErrorCode: errorCode,
				Time: split[1],
			}
		}
	}

	return &TimeResponse{}
}