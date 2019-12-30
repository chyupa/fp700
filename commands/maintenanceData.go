package commands

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/utils"
	"strconv"
	"strings"
)

type MaintenanceDataResponse struct {
	OperName string
	Header1 string
	Header2 string
	Header3 string
	Header4 string
	Header5 string
	Header6 string
	Footer1 string
	Footer2 string
	NZReport string
	Fiscalized string
	ActiveVatGroups string
	NHeaderChanges int
	NVatChanges string
	OperatingMode string
}

func MaintenanceData() MaintenanceDataResponse {
	var mdResponse MaintenanceDataResponse

	operNameResponse, _ := fp700.SendCommand(255, "OperName\t0\t\t")
	operNameDecoded := utils.DecodeMessage(operNameResponse)
	operNameSplit := strings.Split(operNameDecoded, "\t")
	mdResponse.OperName = operNameSplit[1]

	// header
	header1Response, _ := fp700.SendCommand(255, "Header\t0\t\t")
	header1Decoded := utils.DecodeMessage(header1Response)
	header1Split := strings.Split(header1Decoded, "\t")
	mdResponse.Header1 = header1Split[1]

	header2Response, _ := fp700.SendCommand(255, "Header\t1\t\t")
	header2Decoded := utils.DecodeMessage(header2Response)
	header2Split := strings.Split(header2Decoded, "\t")
	mdResponse.Header2 = header2Split[1]

	header3Response, _ := fp700.SendCommand(255, "Header\t2\t\t")
	header3Decoded := utils.DecodeMessage(header3Response)
	header3Split := strings.Split(header3Decoded, "\t")
	mdResponse.Header3 = header3Split[1]

	header4Response, _ := fp700.SendCommand(255, "Header\t3\t\t")
	header4Decoded := utils.DecodeMessage(header4Response)
	header4Split := strings.Split(header4Decoded, "\t")
	mdResponse.Header4 = header4Split[1]

	header5Response, _ := fp700.SendCommand(255, "Header\t4\t\t")
	header5Decoded := utils.DecodeMessage(header5Response)
	header5Split := strings.Split(header5Decoded, "\t")
	mdResponse.Header5 = header5Split[1]

	header6Response, _ := fp700.SendCommand(255, "Header\t5\t\t")
	header6Decoded := utils.DecodeMessage(header6Response)
	header6Split := strings.Split(header6Decoded, "\t")
	mdResponse.Header6= header6Split[1]

	// footer
	footer1Response, _ := fp700.SendCommand(255, "Footer\t0\t\t")
	footer1Decoded := utils.DecodeMessage(footer1Response)
	footer1Split := strings.Split(footer1Decoded, "\t")
	mdResponse.Footer1= footer1Split[1]

	footer2Response, _ := fp700.SendCommand(255, "Footer\t1\t\t")
	footer2Decoded := utils.DecodeMessage(footer2Response)
	footer2Split := strings.Split(footer2Decoded, "\t")
	mdResponse.Footer2= footer2Split[1]

	// nz report
	nzReportResponse, _ := fp700.SendCommand(255, "nZreport\t\t\t")
	nzReportDecoded := utils.DecodeMessage(nzReportResponse)
	nzReportSplit := strings.Split(nzReportDecoded, "\t")
	mdResponse.NZReport= nzReportSplit[1]

	// Fiscalized
	fiscalizedResponse, _ := fp700.SendCommand(255, "Fiscalized\t\t\t")
	fiscalizedDecoded := utils.DecodeMessage(fiscalizedResponse)
	fiscalizedSplit := strings.Split(fiscalizedDecoded, "\t")
	mdResponse.Fiscalized= fiscalizedSplit[1]

	// ActiveVatGroups
	// TODO: simething complicated here; leave it for now
	//activeVatGroupsResponse, _ := fp700.SendCommand(50, "")
	//activeVatGroupsDecoded := utils.DecodeMessage(activeVatGroupsResponse)
	//activeVatGroupsSplit := strings.Split(activeVatGroupsDecoded, "\t")
	//mdResponse.ActiveVatGroups= activeVatGroupsSplit[1]

	// nHeaderChanges
	nHeaderChangesResponse, _ := fp700.SendCommand(43, "I\t")
	nHeaderChangesDecoded := utils.DecodeMessage(nHeaderChangesResponse)
	nHeaderChangesSplit := strings.Split(nHeaderChangesDecoded, "\t")
	maxChanges, _ := strconv.Atoi(nHeaderChangesSplit[2])
	currentChanges, _ := strconv.Atoi(nHeaderChangesSplit[1])
	mdResponse.NHeaderChanges = maxChanges - currentChanges

	// nVatChanges
	nVatChangesResponse, _ := fp700.SendCommand(255, "nVatChanges\t\t\t")
	nVatChangesDecoded := utils.DecodeMessage(nVatChangesResponse)
	nVatChangesSplit := strings.Split(nVatChangesDecoded, "\t")
	mdResponse.NVatChanges = nVatChangesSplit[1]

	// EcrMode
	operatingModeResponse, _ := fp700.SendCommand(255, "EcrMode\t\t\t")
	operatingModeDecoded := utils.DecodeMessage(operatingModeResponse)
	operatingModeSplit := strings.Split(operatingModeDecoded, "\t")
	mdResponse.OperatingMode = operatingModeSplit[1]

	return mdResponse

}
