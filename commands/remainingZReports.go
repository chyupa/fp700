package commands

import (
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
)

func RemainingZReports() (int, error) {
	var decodedMessage = &utils.DecodedMessage{}
	zReports, _ := fp700.SendCommand(68, "")
	zReportsDecoded, err := decodedMessage.DecodeMessage(zReports)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return 0, err
	}

	remainingReports, _ := strconv.Atoi(zReportsDecoded[1])

	return remainingReports, nil
}
