/*******************************************************************************
 *
 *   Copyright (c) 2019 AT&T Intellectual Property.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 *
 *******************************************************************************/
package e2pdus

import (
	"fmt"
	"rsp/enums"
	"strings"
	"testing"
)

/*
 * Create and pack an x2ap setup request.
 * Verify the packed representation matches the want value.
 */
func TestBuildPackedResourceStatusRequest(t *testing.T) {
	var testCases = []struct {
		request   ResourceStatusRequestData
		packedPdu string
	}{
		{
			request: ResourceStatusRequestData{
				CellID:						  "0a0b0c:abcd8000",
				MeasurementID:                15,
				MeasurementID2:               0,
				PartialSuccessAllowed:        true,
				PrbPeriodic:                  true,
				TnlLoadIndPeriodic:           true,
				HwLoadIndPeriodic:            true,
				AbsStatusPeriodic:            true,
				RsrpMeasurementPeriodic:      true,
				CsiPeriodic:                  true,
				PeriodicityMS:                enums.ReportingPeriodicity_one_thousand_ms,
				PeriodicityRsrpMeasurementMS: enums.ReportingPeriodicityRSRPMR_one_hundred_20_ms,
				PeriodicityCsiMS:             enums.ReportingPeriodicityCSIR_ms5,
			},
			packedPdu: "0009003c0000080027000300000e001c00010000260004fe000000001d400d00001f4008000a0b0cabcd8000001e4001000040400100006d4001000091400100",
		},
		{
			request: ResourceStatusRequestData{
				CellID:						  "0a0b0c:abcd8000",
				MeasurementID:                15,
				MeasurementID2:               0,
				PartialSuccessAllowed:        true,
				PrbPeriodic:                  true,
				TnlLoadIndPeriodic:           true,
				HwLoadIndPeriodic:            true,
				AbsStatusPeriodic:            true,
				RsrpMeasurementPeriodic:      true,
				CsiPeriodic:                  true,
				PeriodicityMS:                0,
				PeriodicityRsrpMeasurementMS: 0,
				PeriodicityCsiMS:             0,
			},
			packedPdu: "0009002d0000050027000300000e001c00010000260004fe000000001d400d00001f4008000a0b0cabcd80000040400100",
		},
		{
			request: ResourceStatusRequestData{
				CellID:						  "0a0b0c:abcd8000",
				MeasurementID:                15,
				MeasurementID2:               0,
				PartialSuccessAllowed:        true,
				PrbPeriodic:                  false,
				TnlLoadIndPeriodic:           false,
				HwLoadIndPeriodic:            false,
				AbsStatusPeriodic:            true,
				RsrpMeasurementPeriodic:      true,
				CsiPeriodic:                  true,
				PeriodicityMS:                0,
				PeriodicityRsrpMeasurementMS: 0,
				PeriodicityCsiMS:             0,
			},
			packedPdu: "0009002d0000050027000300000e001c000100002600040e000000001d400d00001f4008000a0b0cabcd80000040400100",
		},
		{
			request: ResourceStatusRequestData{
				CellID:						  "0a0b0c:abcd8000",
				MeasurementID:                15,
				MeasurementID2:               0,
				PartialSuccessAllowed:        false,
				PrbPeriodic:                  false,
				TnlLoadIndPeriodic:           false,
				HwLoadIndPeriodic:            false,
				AbsStatusPeriodic:            true,
				RsrpMeasurementPeriodic:      true,
				CsiPeriodic:                  true,
				PeriodicityMS:                enums.ReportingPeriodicity_ten_thousand_ms,
				PeriodicityRsrpMeasurementMS: enums.ReportingPeriodicityRSRPMR_six_hundred_40_ms,
				PeriodicityCsiMS:             enums.ReportingPeriodicityCSIR_ms80,
			},
			packedPdu: "000900370000070027000300000e001c000100002600040e000000001d400d00001f4008000a0b0cabcd8000001e400160006d4001600091400140",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.packedPdu, func(t *testing.T) {

			payload, asString, err := BuildPackedResourceStatusRequest(enums.Registration_Request_start, &tc.request, MaxAsn1PackedBufferSize /*max packed buffer*/, MaxAsn1CodecMessageBufferSize /*max message buffer*/, true /*withDebug*/)
			if err != nil {
				t.Errorf("want: success, got: pack failed. Error: %v\n", err)
			} else {
				t.Logf("packed resource status request (size=%d): %x\n\n%s", len(payload), payload,asString)
				tmp := fmt.Sprintf("%x", payload)
				if len(tmp) != len(tc.packedPdu) {
					t.Errorf("want packed len:%d, got: %d\n", len(tc.packedPdu)/2, len(payload)/2)
				}

				if strings.Compare(tmp, tc.packedPdu) != 0 {
					t.Errorf("\nwant :\t[%s]\n got: \t\t[%s]\n", tc.packedPdu, tmp)
				}
			}
		})
	}
}

/*Packing error*/

func TestBuildPackedResourceStatusRequestError(t *testing.T) {
	request := ResourceStatusRequestData{
		CellID:						  "0a0b0c:abcd8000",
		MeasurementID:                15,
		MeasurementID2:               0,
		PrbPeriodic:                  true,
		TnlLoadIndPeriodic:           true,
		HwLoadIndPeriodic:            true,
		AbsStatusPeriodic:            true,
		RsrpMeasurementPeriodic:      true,
		CsiPeriodic:                  true,
		PeriodicityMS:                enums.ReportingPeriodicity_one_thousand_ms,
		PeriodicityRsrpMeasurementMS: enums.ReportingPeriodicityRSRPMR_one_hundred_20_ms,
		PeriodicityCsiMS:             enums.ReportingPeriodicityCSIR_ms5,
	}
	expected:= "packing error: #src/asn1codec_utils.c.pack_pdu_aux - Encoded output of E2AP-PDU, is too big"
	_, _, err := BuildPackedResourceStatusRequest(enums.Registration_Request_start, &request, 40 /*max packed buffer*/, MaxAsn1CodecMessageBufferSize /*max message buffer*/, true /*withDebug*/)
	if err != nil {
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("want failure:[%s], got: [%s]\n", expected, err)
		}
	} else {
		t.Errorf("want failure: ...%s..., got: success", expected)

	}
}

func TestBuildPackedResourceStatusInvalidCellID(t *testing.T) {
	request := ResourceStatusRequestData{
		CellID:						  "0a0b0cabcd8000",
		MeasurementID:                15,
		MeasurementID2:               0,
		PrbPeriodic:                  true,
		TnlLoadIndPeriodic:           true,
		HwLoadIndPeriodic:            true,
		AbsStatusPeriodic:            true,
		RsrpMeasurementPeriodic:      true,
		CsiPeriodic:                  true,
		PeriodicityMS:                enums.ReportingPeriodicity_one_thousand_ms,
		PeriodicityRsrpMeasurementMS: enums.ReportingPeriodicityRSRPMR_one_hundred_20_ms,
		PeriodicityCsiMS:             enums.ReportingPeriodicityCSIR_ms5,
	}
	expected:= "BuildPackedResourceStatusRequest() - unexpected CellID value [0a0b0cabcd8000] (want: \"<PLMNIdentifier>:<eUTRANCellIdentifier>\"), err: unexpected EOF"
	_, _, err := BuildPackedResourceStatusRequest(enums.Registration_Request_start, &request, MaxAsn1PackedBufferSize /*max packed buffer*/, MaxAsn1CodecMessageBufferSize /*max message buffer*/, true /*withDebug*/)
	if err != nil {
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("want failure:[%s], got: [%s]\n", expected, err)
		}
	} else {
		t.Errorf("want failure: ...%s..., got: success", expected)

	}
}

func TestBuildPackedResourceStatusInvalidPeriodicity(t *testing.T) {
	request := ResourceStatusRequestData{
		CellID:						  "0a0b0c:abcd8000",
		MeasurementID:                15,
		MeasurementID2:               0,
		PrbPeriodic:                  true,
		TnlLoadIndPeriodic:           true,
		HwLoadIndPeriodic:            true,
		AbsStatusPeriodic:            true,
		RsrpMeasurementPeriodic:      true,
		CsiPeriodic:                  true,
		PeriodicityMS:                22,
		PeriodicityRsrpMeasurementMS: enums.ReportingPeriodicityRSRPMR_one_hundred_20_ms,
		PeriodicityCsiMS:             enums.ReportingPeriodicityCSIR_ms5,
	}
	expected:= "BuildPackedResourceStatusRequest - packing error: #src/asn1codec_utils.c.pack_pdu_aux - Failed to encode E2AP-PDU, error = 9 Bad file descriptor"
	_, _, err := BuildPackedResourceStatusRequest(enums.Registration_Request_start, &request, MaxAsn1PackedBufferSize /*max packed buffer*/, MaxAsn1CodecMessageBufferSize /*max message buffer*/, true /*withDebug*/)
	if err != nil {
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("want failure:[%s], got: [%s]\n", expected, err)
		}
	} else {
		t.Errorf("want failure: ...%s..., got: success", expected)

	}
}
