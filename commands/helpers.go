package commands

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700-server/utils/logger"
	"github.com/chyupa/fp700/utils"
)

type Diagnostic struct {
	Model        string
	Firmware     string
	SerialNumber string
}

type DeviceInformation struct {
	FabricationNumber string
	FiscalNumber      string
	CIF               string
}

type MaintenanceDataResponse struct {
	OperName          string
	CurrNameLocal     string
	OperPasw          string
	Time              string
	Header1           string
	Header2           string
	Header3           string
	Header4           string
	Header5           string
	Header6           string
	Footer1           string
	Footer2           string
	NZReport          string
	Fiscalized        string
	ActiveVatGroups   string
	NHeaderChanges    int
	NVatChanges       string
	OperatingMode     string
	Diagnostic        Diagnostic
	DeviceInformation DeviceInformation
}

func MaintenanceData() (MaintenanceDataResponse, error) {
	var decodedMessage = &utils.DecodedMessage{}
	var mdResponse MaintenanceDataResponse

	operNameResponse, _ := fp700.SendCommand(255, "OperName\t0\t\t")
	operNameDecoded, err := decodedMessage.DecodeMessage(operNameResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.OperName = operNameDecoded[1]

	currNameLocalResponse, _ := fp700.SendCommand(255, "CurrNameLocal\t\t\t")
	currNameLocalDecoded, err := decodedMessage.DecodeMessage(currNameLocalResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.CurrNameLocal = currNameLocalDecoded[1]

	operPassResponse, _ := fp700.SendCommand(255, "OperPasw\t0\t\t")
	operPassDecoded, err := decodedMessage.DecodeMessage(operPassResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.OperPasw = operPassDecoded[1]

	timeResponse, _ := fp700.SendCommand(62, "")
	timeDecoded, err := decodedMessage.DecodeMessage(timeResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Time = timeDecoded[1]

	// header
	header1Response, _ := fp700.SendCommand(255, "Header\t0\t\t")
	header1Decoded, err := decodedMessage.DecodeMessage(header1Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Header1 = header1Decoded[1]

	header2Response, _ := fp700.SendCommand(255, "Header\t1\t\t")
	header2Decoded, err := decodedMessage.DecodeMessage(header2Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Header2 = header2Decoded[1]

	header3Response, _ := fp700.SendCommand(255, "Header\t2\t\t")
	header3Decoded, err := decodedMessage.DecodeMessage(header3Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Header3 = header3Decoded[1]

	header4Response, _ := fp700.SendCommand(255, "Header\t3\t\t")
	header4Decoded, err := decodedMessage.DecodeMessage(header4Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Header4 = header4Decoded[1]

	header5Response, _ := fp700.SendCommand(255, "Header\t4\t\t")
	header5Decoded, err := decodedMessage.DecodeMessage(header5Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Header5 = header5Decoded[1]

	header6Response, _ := fp700.SendCommand(255, "Header\t5\t\t")
	header6Decoded, err := decodedMessage.DecodeMessage(header6Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Header6 = header6Decoded[1]

	// footer
	footer1Response, _ := fp700.SendCommand(255, "Footer\t0\t\t")
	footer1Decoded, err := decodedMessage.DecodeMessage(footer1Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Footer1 = footer1Decoded[1]

	footer2Response, _ := fp700.SendCommand(255, "Footer\t1\t\t")
	footer2Decoded, err := decodedMessage.DecodeMessage(footer2Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Footer2 = footer2Decoded[1]

	// nz report
	nzReportResponse, _ := fp700.SendCommand(255, "nZreport\t\t\t")
	nzReportDecoded, err := decodedMessage.DecodeMessage(nzReportResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.NZReport = nzReportDecoded[1]

	// Fiscalized
	fiscalizedResponse, _ := fp700.SendCommand(255, "Fiscalized\t\t\t")
	fiscalizedDecoded, err := decodedMessage.DecodeMessage(fiscalizedResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.Fiscalized = fiscalizedDecoded[1]

	// ActiveVatGroups
	activeVatGroupsResponse, _ := fp700.SendCommand(50, "")
	activeVatGroupsDecoded, err := decodedMessage.DecodeMessage(activeVatGroupsResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.ActiveVatGroups = activeVatGroupsDecoded[2]

	// nHeaderChanges
	nHeaderChangesResponse, _ := fp700.SendCommand(43, "I\t")
	nHeaderChangesDecoded, err := decodedMessage.DecodeMessage(nHeaderChangesResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	maxChanges, _ := strconv.Atoi(nHeaderChangesDecoded[2])
	currentChanges, _ := strconv.Atoi(nHeaderChangesDecoded[1])
	mdResponse.NHeaderChanges = maxChanges - currentChanges

	// nVatChanges
	nVatChangesResponse, _ := fp700.SendCommand(255, "nVatChanges\t\t\t")
	nVatChangesDecoded, err := decodedMessage.DecodeMessage(nVatChangesResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.NVatChanges = nVatChangesDecoded[1]

	// EcrMode
	operatingModeResponse, _ := fp700.SendCommand(255, "EcrMode\t\t\t")
	operatingModeDecoded, err := decodedMessage.DecodeMessage(operatingModeResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdResponse.OperatingMode = operatingModeDecoded[1]

	// Diagnostic
	diagnosticResponse, _ := fp700.SendCommand(90, "1\t")
	diagnosticResponseDecoded, err := decodedMessage.DecodeMessage(diagnosticResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}
	mdDiagnosticResponse := Diagnostic{
		Model:        diagnosticResponseDecoded[1],
		Firmware:     fmt.Sprintf("Rev %s / %s", diagnosticResponseDecoded[2], diagnosticResponseDecoded[3]),
		SerialNumber: diagnosticResponseDecoded[7],
	}
	mdResponse.Diagnostic = mdDiagnosticResponse

	// Device Information
	deviceInformationResponse, _ := fp700.SendCommand(123, "1\t")
	deviceInformationResponseDecoded, err := decodedMessage.DecodeMessage(deviceInformationResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return mdResponse, err
	}

	mdDeviceInformationResponse := DeviceInformation{
		FabricationNumber: deviceInformationResponseDecoded[1],
		FiscalNumber:      deviceInformationResponseDecoded[2],
		CIF:               strings.Replace(deviceInformationResponseDecoded[5], "CIF: ", "", 1),
	}

	mdResponse.DeviceInformation = mdDeviceInformationResponse

	return mdResponse, nil
}

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

func SetHeadersData(data SetHeadersDataRequest) (SetHeadersDataResponse, error) {
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
		fmt.Println(err)
		logger.Error.Println(err)
		return shdResponse, err
	}

	shdResponse.ErrorCode, _ = strconv.Atoi(readHeaderDecoded[0])
	shdResponse.HdrChanges, _ = strconv.Atoi(readHeaderDecoded[1])
	shdResponse.MaxHdrChanges, _ = strconv.Atoi(readHeaderDecoded[2])
	shdResponse.MaxHdrLines, _ = strconv.Atoi(readHeaderDecoded[3])

	return shdResponse, nil
}

func GetHeaderChanges() (SetHeadersDataResponse, error) {
	var decodedMessage = &utils.DecodedMessage{}
	var shdResponse SetHeadersDataResponse

	// read information regarding header changes
	readHeader, _ := fp700.SendCommand(43, "I\t")
	readHeaderDecoded, err := decodedMessage.DecodeMessage(readHeader)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return shdResponse, err
	}

	shdResponse.ErrorCode, _ = strconv.Atoi(readHeaderDecoded[0])
	shdResponse.HdrChanges, _ = strconv.Atoi(readHeaderDecoded[1])
	shdResponse.MaxHdrChanges, _ = strconv.Atoi(readHeaderDecoded[2])
	shdResponse.MaxHdrLines, _ = strconv.Atoi(readHeaderDecoded[3])

	return shdResponse, nil
}

type FootersDataResponse struct {
	Footer1 string
	Footer2 string
}

func FootersData() (FootersDataResponse, error) {
	var decodedMessage = &utils.DecodedMessage{}
	var fdResponse FootersDataResponse

	// header
	footer1Response, _ := fp700.SendCommand(255, "Footer\t0\t\t")
	footer1Decoded, err := decodedMessage.DecodeMessage(footer1Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return fdResponse, err
	}
	fdResponse.Footer1 = footer1Decoded[1]

	footer2Response, _ := fp700.SendCommand(255, "Footer\t1\t\t")
	footer2Decoded, err := decodedMessage.DecodeMessage(footer2Response)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return fdResponse, err
	}
	fdResponse.Footer2 = footer2Decoded[1]

	return fdResponse, nil
}

type SetFootersDataRequest struct {
	TopLine    string `json:"topLineValue"`
	BottomLine string `json:"bottomLineValue"`
}

func SetFootersData(data SetFootersDataRequest) error {
	var decodedMessage = &utils.DecodedMessage{}
	topFooter, _ := fp700.SendCommand(255, "Footer\t0\t"+data.TopLine+"\t")
	_, err := decodedMessage.DecodeMessage(topFooter)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return err
	}
	bottomFooter, _ := fp700.SendCommand(255, "Footer\t1\t"+data.BottomLine+"\t")
	_, err = decodedMessage.DecodeMessage(bottomFooter)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return err
	}

	return nil
}

func PrintDiagnostic() {
	fp700.SendCommand(71, "0\t")
}

func Time() (string, error) {
	var decodedMessage = &utils.DecodedMessage{}
	reply, err := fp700.SendCommand(62, "")
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return "", err
	}

	if len(reply) > 2 {
		msg, err := decodedMessage.DecodeMessage(reply)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return msg[1], nil
	}

	return "", errors.New("something went wrong")
}

type SetTimeRequest struct {
	Time string `json:"time"`
}

func SetTime(data SetTimeRequest) string {
	var decodedMessage = &utils.DecodedMessage{}

	setTimeResponse, _ := fp700.SendCommand(61, data.Time+"\t")
	setTimeDecoded, err := decodedMessage.DecodeMessage(setTimeResponse)
	if err != nil {
		log.Println(err)
	}

	return setTimeDecoded[0]
}

type InitDisplayRequest struct {
	FirstLine  string `json:"firstLine"'`
	SecondLine string `json:"secondLine"'`
}

func InitDisplay(idRequest InitDisplayRequest) {

	// reset display
	fp700.SendCommand(33, "")

	// set first line
	fp700.SendCommand(47, idRequest.FirstLine+"\t")

	// set second line
	fp700.SendCommand(35, idRequest.SecondLine+"\t")
}

func FabricationNumber() {
	// print fabrication number; SAM Module
	fp700.SendCommand(71, "5\t")
}

func GetOperatorName() (string, error) {

	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, "OperName\t0\t\t")

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[1], nil
}

type SetOperatorNameRequest struct {
	Name string `json:"name"`
}

func SetOperatorName(data SetOperatorNameRequest) (string, error) {
	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, fmt.Sprintf("OperName\t0\t%s\t", data.Name))

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[0], nil
}

func GetOperatorPassword() (string, error) {

	var decodedMessage utils.DecodedMessage
	cmd, _ := fp700.SendCommand(255, "OperPasw\t0\t\t")

	response, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return "", err
	}

	return response[1], nil
}

type SetOperatorPasswordRequest struct {
	OldPass string `json:"oldPass"`
	NewPass string `json:"newPass"`
}

func SetOperatorPassword(data SetOperatorPasswordRequest) error {
	var decodedMessage utils.DecodedMessage
	payload := fmt.Sprintf("1\t%s\t%s\t", data.OldPass, data.NewPass)
	cmd, _ := fp700.SendCommand(101, payload)

	_, err := decodedMessage.DecodeMessage(cmd)
	if err != nil {
		return err
	}

	return nil
}

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

func GetVat() (string, error) {
	var decodedMessage = &utils.DecodedMessage{}

	// ActiveVatGroups
	activeVatGroupsResponse, _ := fp700.SendCommand(50, "")
	activeVatGroupsDecoded, err := decodedMessage.DecodeMessage(activeVatGroupsResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return "", err
	}
	return activeVatGroupsDecoded[2], nil
}

func GetVatChanges() (int, error) {
	var decodedMessage = &utils.DecodedMessage{}

	// nVatChanges
	nVatChangesResponse, _ := fp700.SendCommand(255, "nVatChanges\t\t\t")
	nVatChangesDecoded, err := decodedMessage.DecodeMessage(nVatChangesResponse)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return -1, nil
	}
	vatChanges, _ := strconv.Atoi(nVatChangesDecoded[1])
	return vatChanges, nil
}

type SetVatRequest struct {
	Vat string `json:"vat"`
}

func SetVat(data SetVatRequest) (int, error) {
	var decodedMessage = &utils.DecodedMessage{}

	setVatResponse, _ := fp700.SendCommand(83, data.Vat+"\t")
	setVatDecoded, err := decodedMessage.DecodeMessage(setVatResponse)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	remainingChanges, err := strconv.Atoi(setVatDecoded[1])
	return remainingChanges, nil
}

type CloseDmjeRequest struct {
	servicePass string
}

func CloseDmje(data CloseDmjeRequest) (int, error) {
	fp700.SendCommand(253, "0\t"+data.servicePass+"\t")
	closeResponse, _ := fp700.SendCommand(253, "2\t"+data.servicePass+"\t\t")
	var decodedMessage = &utils.DecodedMessage{}
	msg, err := decodedMessage.DecodeMessage(closeResponse)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	errorCode, _ := strconv.Atoi(msg[0])
	return errorCode, nil
}

type FiscalizeRequest struct {
	date        string `json:"dateTime"`
	fabrication string `json:"fabricationNumber"`
	vat         string `json:"vatValue"`
	tax         string `json:"taxNo"`
	cif         string `json:"cif"`
}

type FiscalizeResponse struct {
	ErrorCode int
}

func Fiscalize(data FiscalizeRequest) (FiscalizeResponse, error) {
	var decodedMessage = &utils.DecodedMessage{}
	var fResponse FiscalizeResponse
	// set time
	timeResponse, _ := fp700.SendCommand(61, data.date+"\t")
	log.Println(timeResponse)

	// set serial number
	serialNumberResponse, _ := fp700.SendCommand(91, data.fabrication+"\t")
	log.Println(serialNumberResponse)

	// set VAT number
	vatResponse, _ := fp700.SendCommand(83, data.vat+"\t")
	log.Println(vatResponse)

	// set Fiscalization
	fiscalizeResponse, _ := fp700.SendCommand(72, data.tax+"\t"+data.cif+"\t")
	log.Println(fiscalizeResponse)
	msg, err := decodedMessage.DecodeMessage(fiscalizeResponse)
	if err != nil {
		log.Println(err)
		return FiscalizeResponse{}, err
	}

	fResponse.ErrorCode, _ = strconv.Atoi(msg[0])

	// disable close open receipt printer option
	fp700.SendCommand(255, "DsblKeyCloseReceipt\t\t1\t")
	// disable cancel receipt printer option
	fp700.SendCommand(255, "DsblKeyCancelReceipt\t\t1\t")
	// disable generation of fiscal memory reports
	fp700.SendCommand(255, "DsblKeyFmReports\t\t1\t")

	return fResponse, nil
}
