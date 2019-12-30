package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
)
type HeadersDataResponse struct {
	Header1 string
	Header2 string
	Header3 string
	Header4 string
	Header5 string
	Header6 string
}

func HeadersData() HeadersDataResponse {
	var hdResponse HeadersDataResponse

	// header
	header1Response, _ := fp700.SendCommand(255, "Header\t0\t\t")
	header1Decoded := utils.DecodeMessage(header1Response)
	header1Split := strings.Split(header1Decoded, "\t")
	hdResponse.Header1 = header1Split[1]

	header2Response, _ := fp700.SendCommand(255, "Header\t1\t\t")
	header2Decoded := utils.DecodeMessage(header2Response)
	header2Split := strings.Split(header2Decoded, "\t")
	hdResponse.Header2 = header2Split[1]

	header3Response, _ := fp700.SendCommand(255, "Header\t2\t\t")
	header3Decoded := utils.DecodeMessage(header3Response)
	header3Split := strings.Split(header3Decoded, "\t")
	hdResponse.Header3 = header3Split[1]

	header4Response, _ := fp700.SendCommand(255, "Header\t3\t\t")
	header4Decoded := utils.DecodeMessage(header4Response)
	header4Split := strings.Split(header4Decoded, "\t")
	hdResponse.Header4 = header4Split[1]

	header5Response, _ := fp700.SendCommand(255, "Header\t4\t\t")
	header5Decoded := utils.DecodeMessage(header5Response)
	header5Split := strings.Split(header5Decoded, "\t")
	hdResponse.Header5 = header5Split[1]

	header6Response, _ := fp700.SendCommand(255, "Header\t5\t\t")
	header6Decoded := utils.DecodeMessage(header6Response)
	header6Split := strings.Split(header6Decoded, "\t")
	hdResponse.Header6 = header6Split[1]

	return hdResponse
}

type SetHeadersDataResponse struct {
	ErrorCode int
	HdrChanges int
	MaxHdrChanges int
	MaxHdrLines int
}

type SetHeadersDataRequest struct {
	Header1 string `json:"headerLine1"`
	Header2 string `json:"headerLine2"`
	Header3 string `json:"headerLine3"`
	Header4 string `json:"headerLine4"`
	Header5 string `json:"headerLine5"`
	Header6 string `json:"headerLine6"`
}

func SetHeadersData(data SetHeadersDataRequest) SetHeadersDataResponse {
	var shdResponse SetHeadersDataResponse
	fp700.SendCommand(43, "W\t1\t" + data.Header1 + "\t")
	fp700.SendCommand(43, "W\t2\t" + data.Header2 + "\t")
	fp700.SendCommand(43, "W\t3\t" + data.Header3 + "\t")
	fp700.SendCommand(43, "W\t4\t" + data.Header4 + "\t")
	fp700.SendCommand(43, "W\t5\t" + data.Header5 + "\t")
	fp700.SendCommand(43, "W\t6\t" + data.Header6 + "\t")

	// push changes to fiscal memory
	fp700.SendCommand(43, "W\t")

	// read information regarding header changes
	readHeader, _ := fp700.SendCommand(43, "I\t")
	readHeaderDecoded := utils.DecodeMessage(readHeader)
	split := strings.Split(readHeaderDecoded, "\t")

	shdResponse.ErrorCode, _ = strconv.Atoi(split[0])
	shdResponse.HdrChanges, _ = strconv.Atoi(split[1])
	shdResponse.MaxHdrChanges, _ = strconv.Atoi(split[2])
	shdResponse.MaxHdrLines, _ = strconv.Atoi(split[3])

	return shdResponse
}
