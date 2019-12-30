package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
	"time"
)

type MfRequest struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	ReportType string `json:"reportType"`
	ByDate     bool   `json:"byDate"`
}

func ReadMf(mfReq MfRequest) (string, error) {
	var decodedMessage = &utils.DecodedMessage{}

	var command = 95
	payload := fmt.Sprintf("%s\t%s\t%s\t", mfReq.ReportType, mfReq.Start, mfReq.End)
	if mfReq.ByDate {
		command = 94
	}

	// initialize reading from EJ
	initialCommandResponse, err := fp700.SendCommand(command, payload)
	if err != nil {
		log.Println(err)
	}

	_, err = decodedMessage.DecodeMessage(initialCommandResponse)

	if err != nil {
		return "", err
	}

	port, _ := fp700.OpenPort()

	defer port.Close()

	reader := bufio.NewReader(port)

	response := bytes.Buffer{}

	for {
		payload = fmt.Sprintf("3\t%s\t%s\t", mfReq.Start, mfReq.End)
		readCommand := utils.EncodeMessage(command, payload)

		port.Write(readCommand)

		time.Sleep(time.Millisecond * 100)
		reply, _ := reader.ReadBytes(0x03)

		if len(reply) > 1 {
			msg, err := decodedMessage.DecodeMessage(reply)
			if err != nil {
				return "", err
			}

			if len(msg) > 2 {
				response.WriteString(msg[1] + "\r\n")
			} else {
				exitCode, _ := strconv.Atoi(msg[0])
				if exitCode < -1 {
					break
				}
			}
		}

	}

	return response.String(), nil
}
