// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BpStatus byte

const (
	BpStatusNoError                 BpStatus = 0
	BpStatusErrorIncompleteData     BpStatus = 1
	BpStatusErrorNoMeasurement      BpStatus = 2
	BpStatusErrorDataOutOfRange     BpStatus = 3
	BpStatusErrorIrregularHeartRate BpStatus = 4
	BpStatusInvalid                 BpStatus = 0xFF
)

func (b BpStatus) Byte() byte { return byte(b) }

func (b BpStatus) String() string {
	switch b {
	case BpStatusNoError:
		return "no_error"
	case BpStatusErrorIncompleteData:
		return "error_incomplete_data"
	case BpStatusErrorNoMeasurement:
		return "error_no_measurement"
	case BpStatusErrorDataOutOfRange:
		return "error_data_out_of_range"
	case BpStatusErrorIrregularHeartRate:
		return "error_irregular_heart_rate"
	default:
		return "BpStatusInvalid(" + strconv.Itoa(int(b)) + ")"
	}
}

// FromString parse string into BpStatus constant it's represent, return BpStatusInvalid if not found.
func BpStatusFromString(s string) BpStatus {
	switch s {
	case "no_error":
		return BpStatusNoError
	case "error_incomplete_data":
		return BpStatusErrorIncompleteData
	case "error_no_measurement":
		return BpStatusErrorNoMeasurement
	case "error_data_out_of_range":
		return BpStatusErrorDataOutOfRange
	case "error_irregular_heart_rate":
		return BpStatusErrorIrregularHeartRate
	default:
		return BpStatusInvalid
	}
}

// List returns all constants.
func ListBpStatus() []BpStatus {
	return []BpStatus{
		BpStatusNoError,
		BpStatusErrorIncompleteData,
		BpStatusErrorNoMeasurement,
		BpStatusErrorDataOutOfRange,
		BpStatusErrorIrregularHeartRate,
	}
}