package commands

import (
	"bufio"
	"bytes"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

func ReadEj() string {

	// initialize reading from EJ
	fp700.SendCommand(125, "20\t2\t2\t")

	port, _ := fp700.OpenPort()

	reader := bufio.NewReader(port)

	response := bytes.Buffer{}

	for {
		readCommand := utils.EncodeMessage(125, "21\t\t\t")

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

	return response.String();
}
