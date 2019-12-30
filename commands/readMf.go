package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

type MfRequest struct {
	Start string `json:"start"`
	End string `json:"end"`
	ReportType string `json:"reportType"`
	ByDate bool `json:"byDate"`
}

func ReadMf(start string, end string, reportType string, byDate bool) string {
	var command = 95
	if byDate {
		command = 94
	}

	payload := fmt.Sprintf("%s\t%s\t%s\t", reportType, start, end)

	// initialize reading from EJ
	fp700.SendCommand(command, payload)

	port, _ := fp700.OpenPort()

	defer port.Close()

	reader := bufio.NewReader(port)

	response := bytes.Buffer{}

	for {
		payload = fmt.Sprintf("3\t%s\t%s\t", start, end)
		readCommand := utils.EncodeMessage(command, payload)

		port.Write(readCommand)

		time.Sleep(time.Millisecond * 100)
		reply, _ := reader.ReadBytes(0x03)

		if len(reply) > 1 {
			msg := utils.DecodeMessage(reply)
			log.Print(msg)

			split := strings.Split(msg, "\t")
			if len(split) > 2 {
				response.WriteString(split[1] + "\r\n")
			} else {
				exitCode, _ := strconv.Atoi(split[0])
				if exitCode < -1 {
					break
				}
			}
		}

	}

	return response.String()
}
