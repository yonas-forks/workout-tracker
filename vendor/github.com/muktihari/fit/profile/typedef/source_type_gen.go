// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SourceType byte

const (
	SourceTypeAnt                SourceType = 0 // External device connected with ANT
	SourceTypeAntplus            SourceType = 1 // External device connected with ANT+
	SourceTypeBluetooth          SourceType = 2 // External device connected with BT
	SourceTypeBluetoothLowEnergy SourceType = 3 // External device connected with BLE
	SourceTypeWifi               SourceType = 4 // External device connected with Wifi
	SourceTypeLocal              SourceType = 5 // Onboard device
	SourceTypeInvalid            SourceType = 0xFF
)

func (s SourceType) Byte() byte { return byte(s) }

func (s SourceType) String() string {
	switch s {
	case SourceTypeAnt:
		return "ant"
	case SourceTypeAntplus:
		return "antplus"
	case SourceTypeBluetooth:
		return "bluetooth"
	case SourceTypeBluetoothLowEnergy:
		return "bluetooth_low_energy"
	case SourceTypeWifi:
		return "wifi"
	case SourceTypeLocal:
		return "local"
	default:
		return "SourceTypeInvalid(" + strconv.Itoa(int(s)) + ")"
	}
}

// FromString parse string into SourceType constant it's represent, return SourceTypeInvalid if not found.
func SourceTypeFromString(s string) SourceType {
	switch s {
	case "ant":
		return SourceTypeAnt
	case "antplus":
		return SourceTypeAntplus
	case "bluetooth":
		return SourceTypeBluetooth
	case "bluetooth_low_energy":
		return SourceTypeBluetoothLowEnergy
	case "wifi":
		return SourceTypeWifi
	case "local":
		return SourceTypeLocal
	default:
		return SourceTypeInvalid
	}
}

// List returns all constants.
func ListSourceType() []SourceType {
	return []SourceType{
		SourceTypeAnt,
		SourceTypeAntplus,
		SourceTypeBluetooth,
		SourceTypeBluetoothLowEnergy,
		SourceTypeWifi,
		SourceTypeLocal,
	}
}