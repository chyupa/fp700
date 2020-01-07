package commands

import (
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
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
