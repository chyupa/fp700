package commands

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700-server/utils/logger"
	"github.com/chyupa/fp700/utils"
)

type PrintEjRequest struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	ByDate bool   `json:"byDate"`
}

func PrintEj(pEjReq PrintEjRequest) error {
	command := 125
	payload := fmt.Sprintf("13\t%s\t%s\t", pEjReq.Start, pEjReq.End)
	if pEjReq.ByDate {
		command = 124
		payload = fmt.Sprintf("%s\t%s\t2\t", pEjReq.Start, pEjReq.End)

		initialCommand, err := fp700.SendCommand(command, payload)
		if err != nil {
			log.Println(err)
		}

		initialCommandResponse, err := decodedMessage.DecodeMessage(initialCommand)
		if err != nil {
			return err
		}

		pEjReq.ByDate = false
		pEjReq.Start = initialCommandResponse[4]
		pEjReq.End = initialCommandResponse[6]

		return PrintEj(pEjReq)
	}

	_, err := fp700.SendCommand(command, payload)
	fp700.SendCommand(46, "")
	return err
}

type PrintMfRequest struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	ReportType string `json:"reportType"`
	ByDate     bool   `json:"byDate"`
}

func PrintMf(mfReq PrintMfRequest) error {

	var command = 95
	payload := fmt.Sprintf("%s\t%s\t%s\t", mfReq.ReportType, mfReq.Start, mfReq.End)
	if mfReq.ByDate {
		command = 94
	}

	// initialize reading from EJ
	_, err := fp700.SendCommand(command, payload)
	if err != nil {
		return err
	}

	return nil
}

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

type MfRequest struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	ReportType string `json:"reportType"`
	ByDate     bool   `json:"byDate"`
}

func ReadMf(mfReq MfRequest) (string, error) {
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

type File struct {
	XMLName xml.Name `xml:"file"`
	Name    string   `xml:"name,attr"`
	SaveAs  string   `xml:"saveas,attr"`
}
type AnafDir struct {
	XMLName xml.Name `xml:"dir"`
	Files   []File   `xml:"file"`
}
type AnafXml struct {
	XMLName xml.Name `xml:"list"`
	Dir     AnafDir  `xml:"dir"`
}

type AnafFilesResponse struct{}

type AnafFilesRequest struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func AnafFiles(data AnafFilesRequest) (map[string]string, error) {
	initialCommand, err := fp700.SendCommand(128, "0\t"+data.Start+"\t"+data.End+"\t")
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	_, err = decodedMessage.DecodeMessage(initialCommand)
	if err != nil {
		return nil, err
	}

	port, _ := fp700.OpenPort()

	defer port.Close()

	reader := bufio.NewReader(port)

	response := bytes.Buffer{}

	for {
		readCommand := utils.EncodeMessage(128, "1\t")

		port.Write(readCommand)

		time.Sleep(time.Millisecond * 100)
		reply, _ := reader.ReadBytes(0x03)

		if len(reply) > 1 {
			msg, err := decodedMessage.DecodeMessage(reply)
			if err != nil {
				fmt.Println(err)
				logger.Error.Println(err)
				return nil, err
			}
			if len(msg) > 2 && !strings.Contains(msg[1], "XML") {
				response.WriteString(msg[1] + "\r\n")
			} else {
				exitCode, _ := strconv.Atoi(msg[0])
				if exitCode < -1 {
					break
				}
			}
		}

	}

	xmlString, _ := base64.StdEncoding.DecodeString(response.String())
	var anafXml AnafXml

	_ = xml.Unmarshal(xmlString, &anafXml)

	m := map[string]string{}
	for _, fileName := range anafXml.Dir.Files {
		readCommand := utils.EncodeMessage(128, "0\t"+fileName.Name+"\t")
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
				msg, err := decodedMessage.DecodeMessage(reply)
				if err != nil {
					fmt.Println(err)
					logger.Error.Println(err)
					return nil, err
				}
				if len(msg) > 2 && !strings.Contains(msg[1], "xml") {
					response.WriteString(msg[1])
				} else {
					exitCode, _ := strconv.Atoi(msg[0])
					if exitCode < -1 {
						break
					}
				}
			}
		}

		m[fileName.SaveAs] = response.String()
	}

	return m, nil
}

type ReportXResponse struct {
	ErrorCode  int
	nRep       int
	TotX       int
	TotEXEPTAT int
	TotSInv    int
	VatSInv    int
}

func PrintReport(reportType string) (ReportXResponse, error) {
	response := ReportXResponse{}

	payload := fmt.Sprintf("%s\t", reportType)
	reply, err := fp700.SendCommand(69, payload)

	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	if len(reply) > 2 {
		msg, err := decodedMessage.DecodeMessage(reply)
		if err != nil {
			fmt.Println(err)
			logger.Error.Println(err)
			return response, err
		}

		response.ErrorCode, _ = strconv.Atoi(msg[0])
		response.nRep, _ = strconv.Atoi(msg[1])
		response.TotX, _ = strconv.Atoi(msg[2])
		response.TotEXEPTAT, _ = strconv.Atoi(msg[3])
		response.TotSInv, _ = strconv.Atoi(msg[4])
		response.VatSInv, _ = strconv.Atoi(msg[5])
	}

	return response, nil
}
