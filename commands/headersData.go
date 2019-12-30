package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"log"
	"strconv"
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
	var decodedMessage = &utils.DecodedMessage{}
	var hdResponse HeadersDataResponse

	// header
	header1Response, _ := fp700.SendCommand(255, "Header\t0\t\t")
	header1Decoded, err := decodedMessage.DecodeMessage(header1Response)
	if err != nil {
		log.Println(err)
	}
	hdResponse.Header1 = header1Decoded[1]

	header2Response, _ := fp700.SendCommand(255, "Header\t1\t\t")
	header2Decoded, err := decodedMessage.DecodeMessage(header2Response)
	if err != nil {
		log.Println(err)
	}
	hdResponse.Header2 = header2Decoded[1]

	header3Response, _ := fp700.SendCommand(255, "Header\t2\t\t")
	header3Decoded, err := decodedMessage.DecodeMessage(header3Response)
	if err != nil {
		log.Println(err)
	}
	hdResponse.Header3 = header3Decoded[1]

	header4Response, _ := fp700.SendCommand(255, "Header\t3\t\t")
	header4Decoded, err := decodedMessage.DecodeMessage(header4Response)
	if err != nil {
		log.Println(err)
	}
	hdResponse.Header4 = header4Decoded[1]

	header5Response, _ := fp700.SendCommand(255, "Header\t4\t\t")
	header5Decoded, err := decodedMessage.DecodeMessage(header5Response)
	if err != nil {
		log.Println(err)
	}
	hdResponse.Header5 = header5Decoded[1]

	header6Response, _ := fp700.SendCommand(255, "Header\t5\t\t")
	header6Decoded, err := decodedMessage.DecodeMessage(header6Response)
	if err != nil {
		log.Println(err)
	}
	hdResponse.Header6 = header6Decoded[1]

	return hdResponse
}

type SetHeadersDataResponse struct {
	ErrorCode     int
	HdrChanges    int
	MaxHdrChanges int
	MaxHdrLines   int
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
	var decodedMessage = &utils.DecodedMessage{}
	var shdResponse SetHeadersDataResponse
	fp700.SendCommand(43, "W\t1\t"+data.Header1+"\t")
	fp700.SendCommand(43, "W\t2\t"+data.Header2+"\t")
	fp700.SendCommand(43, "W\t3\t"+data.Header3+"\t")
	fp700.SendCommand(43, "W\t4\t"+data.Header4+"\t")
	fp700.SendCommand(43, "W\t5\t"+data.Header5+"\t")
	fp700.SendCommand(43, "W\t6\t"+data.Header6+"\t")

	// push changes to fiscal memory
	fp700.SendCommand(43, "W\t")

	// read information regarding header changes
	readHeader, _ := fp700.SendCommand(43, "I\t")
	readHeaderDecoded, err := decodedMessage.DecodeMessage(readHeader)
	if err != nil {
		log.Println(err)
	}

	shdResponse.ErrorCode, _ = strconv.Atoi(readHeaderDecoded[0])
	shdResponse.HdrChanges, _ = strconv.Atoi(readHeaderDecoded[1])
	shdResponse.MaxHdrChanges, _ = strconv.Atoi(readHeaderDecoded[2])
	shdResponse.MaxHdrLines, _ = strconv.Atoi(readHeaderDecoded[3])

	return shdResponse
}
