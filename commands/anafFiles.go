package commands

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
	"time"
)
type File struct {
	XMLName xml.Name `xml:"file"`
	Name string `xml:"name,attr"`
	SaveAs string `xml:"saveas,attr"`
}
type AnafDir struct {
	XMLName xml.Name `xml:"dir"`
	Files []File `xml:"file"`
}
type AnafXml struct {
	XMLName xml.Name `xml:"list"`
	Dir AnafDir `xml:"dir"`
}

type AnafFilesResponse struct {

}
type AnafFilesRequest struct {
	Start string `json:"start"`
	End string `json:"end"`
}

func AnafFiles(data AnafFilesRequest) (map[string] string, error) {
	r, e := fp700.SendCommand(128, "0\t" + data.Start + "\t" + data.End + "\t")
	fmt.Println(utils.DecodeMessage(r), e)
	port, _ := fp700.OpenPort()

	defer port.Close()

	reader := bufio.NewReader(port)

	response := bytes.Buffer{}

	for {
		readCommand := utils.EncodeMessage(128, "1\t")

		port.Write(readCommand)

		time.Sleep(time.Millisecond * 100)
		reply, _ := reader.ReadBytes(0x03)

		startingIndex := utils.IndexOf(reply, 1);
		if startingIndex != -1 {
			reply = append(reply[startingIndex:])
		}

		if len(reply) > 1 {
			msg := utils.DecodeMessage(reply)

			split := strings.Split(msg, "\t")
			if len(split) > 2 && !strings.Contains(split[1], "XML") {
				response.WriteString(split[1] + "\r\n")
			} else {
				exitCode, _ := strconv.Atoi(split[0])
				if exitCode < -1 {
					break
				}
			}
		}

	}

	xmlString, _ := base64.StdEncoding.DecodeString(response.String())
	var anafXml AnafXml

	_ = xml.Unmarshal(xmlString, &anafXml)

	m := map[string] string{}
	for _, fileName := range anafXml.Dir.Files {
		readCommand := utils.EncodeMessage(128, "0\t" + fileName.Name + "\t")
		port.Write(readCommand)

		response := bytes.Buffer{}
		for {
			readCommand := utils.EncodeMessage(128, "1\t")

			port.Write(readCommand)

			//time.Sleep(time.Millisecond * 100)
			reply, _ := reader.ReadBytes(0x03)

			startingIndex := utils.IndexOf(reply, 1)
			if startingIndex != -1 {
				reply = append(reply[startingIndex:])
			}

			if len(reply) > 1 {
				msg := utils.DecodeMessage(reply)
				split := strings.Split(msg, "\t")
				if len(split) > 2 && !strings.Contains(split[1], "xml") {
					response.WriteString(split[1])
				} else {
					exitCode, _ := strconv.Atoi(split[0])
					if exitCode < -1 {
						break
						//return m, errors.New(string(exitCode))
					}
				}
			}
		}

		m[fileName.SaveAs] = response.String()
	}

	return m, nil
}
