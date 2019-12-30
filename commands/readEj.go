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

type EjRequest struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	ReportType string `json:"reportType"`
	ByDate     bool   `json:"byDate"`
}

func ReadEj(ejReq EjRequest) (string, error) {
	command := 125
	payload := fmt.Sprintf("%s\t%s\t%s\t", ejReq.ReportType, ejReq.Start, ejReq.End)
	if ejReq.ByDate {
		command = 124
		payload = fmt.Sprintf("%s\t%s\t%s\t", ejReq.Start, ejReq.End, ejReq.ReportType)
	}

	var decodedMessage = &utils.DecodedMessage{}

	// initialize reading from EJ
	initialCommand, err := fp700.SendCommand(command, payload)
	if err != nil {
		log.Println(err)
	}

	initialCommandResponse, err := decodedMessage.DecodeMessage(initialCommand)
	if err != nil {
		return "", err
	}

	if ejReq.ByDate {
		ejReq.ReportType = "20"
		ejReq.Start = initialCommandResponse[4]
		ejReq.End = initialCommandResponse[6]
		ejReq.ByDate = false

		return ReadEj(ejReq)
	}

	port, _ := fp700.OpenPort()

	defer port.Close()

	reader := bufio.NewReader(port)

	response := bytes.Buffer{}

	for {
		readCommand := utils.EncodeMessage(125, "21\t\t\t")

		port.Write(readCommand)

		time.Sleep(time.Millisecond * 100)
		reply, _ := reader.ReadBytes(0x03)

		if len(reply) > 1 {
			msg, err := decodedMessage.DecodeMessage(reply)
			if err != nil {
				log.Println(err)
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
