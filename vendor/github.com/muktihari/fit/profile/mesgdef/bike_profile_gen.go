// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// BikeProfile is a BikeProfile message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type BikeProfile struct {
	Name                     string
	FrontGear                []uint8 // Base: uint8z; Array: [N]; Number of teeth on each gear 0 is innermost
	RearGear                 []uint8 // Base: uint8z; Array: [N]; Number of teeth on each gear 0 is innermost
	Odometer                 uint32  // Scale: 100; Units: m
	MessageIndex             typedef.MessageIndex
	BikeSpdAntId             uint16 // Base: uint16z
	BikeCadAntId             uint16 // Base: uint16z
	BikeSpdcadAntId          uint16 // Base: uint16z
	BikePowerAntId           uint16 // Base: uint16z
	CustomWheelsize          uint16 // Scale: 1000; Units: m
	AutoWheelsize            uint16 // Scale: 1000; Units: m
	BikeWeight               uint16 // Scale: 10; Units: kg
	PowerCalFactor           uint16 // Scale: 10; Units: %
	Sport                    typedef.Sport
	SubSport                 typedef.SubSport
	AutoWheelCal             typedef.Bool
	AutoPowerZero            typedef.Bool
	Id                       uint8
	SpdEnabled               typedef.Bool
	CadEnabled               typedef.Bool
	SpdcadEnabled            typedef.Bool
	PowerEnabled             typedef.Bool
	CrankLength              uint8 // Scale: 2; Offset: -110; Units: mm
	Enabled                  typedef.Bool
	BikeSpdAntIdTransType    uint8 // Base: uint8z
	BikeCadAntIdTransType    uint8 // Base: uint8z
	BikeSpdcadAntIdTransType uint8 // Base: uint8z
	BikePowerAntIdTransType  uint8 // Base: uint8z
	OdometerRollover         uint8 // Rollover counter that can be used to extend the odometer
	FrontGearNum             uint8 // Base: uint8z; Number of front gears
	RearGearNum              uint8 // Base: uint8z; Number of rear gears
	ShimanoDi2Enabled        typedef.Bool

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewBikeProfile creates new BikeProfile struct based on given mesg.
// If mesg is nil, it will return BikeProfile with all fields being set to its corresponding invalid value.
func NewBikeProfile(mesg *proto.Message) *BikeProfile {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &BikeProfile{
		MessageIndex:             typedef.MessageIndex(vals[254].Uint16()),
		Name:                     vals[0].String(),
		Sport:                    typedef.Sport(vals[1].Uint8()),
		SubSport:                 typedef.SubSport(vals[2].Uint8()),
		Odometer:                 vals[3].Uint32(),
		BikeSpdAntId:             vals[4].Uint16z(),
		BikeCadAntId:             vals[5].Uint16z(),
		BikeSpdcadAntId:          vals[6].Uint16z(),
		BikePowerAntId:           vals[7].Uint16z(),
		CustomWheelsize:          vals[8].Uint16(),
		AutoWheelsize:            vals[9].Uint16(),
		BikeWeight:               vals[10].Uint16(),
		PowerCalFactor:           vals[11].Uint16(),
		AutoWheelCal:             vals[12].Bool(),
		AutoPowerZero:            vals[13].Bool(),
		Id:                       vals[14].Uint8(),
		SpdEnabled:               vals[15].Bool(),
		CadEnabled:               vals[16].Bool(),
		SpdcadEnabled:            vals[17].Bool(),
		PowerEnabled:             vals[18].Bool(),
		CrankLength:              vals[19].Uint8(),
		Enabled:                  vals[20].Bool(),
		BikeSpdAntIdTransType:    vals[21].Uint8z(),
		BikeCadAntIdTransType:    vals[22].Uint8z(),
		BikeSpdcadAntIdTransType: vals[23].Uint8z(),
		BikePowerAntIdTransType:  vals[24].Uint8z(),
		OdometerRollover:         vals[37].Uint8(),
		FrontGearNum:             vals[38].Uint8z(),
		FrontGear:                vals[39].SliceUint8(),
		RearGearNum:              vals[40].Uint8z(),
		RearGear:                 vals[41].SliceUint8(),
		ShimanoDi2Enabled:        vals[44].Bool(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts BikeProfile into proto.Message. If options is nil, default options will be used.
func (m *BikeProfile) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumBikeProfile}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if m.Sport != typedef.SportInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if m.SubSport != typedef.SubSportInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.SubSport))
		fields = append(fields, field)
	}
	if m.Odometer != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.Odometer)
		fields = append(fields, field)
	}
	if m.BikeSpdAntId != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.BikeSpdAntId)
		fields = append(fields, field)
	}
	if m.BikeCadAntId != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.BikeCadAntId)
		fields = append(fields, field)
	}
	if m.BikeSpdcadAntId != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint16(m.BikeSpdcadAntId)
		fields = append(fields, field)
	}
	if m.BikePowerAntId != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint16(m.BikePowerAntId)
		fields = append(fields, field)
	}
	if m.CustomWheelsize != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint16(m.CustomWheelsize)
		fields = append(fields, field)
	}
	if m.AutoWheelsize != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(m.AutoWheelsize)
		fields = append(fields, field)
	}
	if m.BikeWeight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint16(m.BikeWeight)
		fields = append(fields, field)
	}
	if m.PowerCalFactor != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint16(m.PowerCalFactor)
		fields = append(fields, field)
	}
	if m.AutoWheelCal < 2 {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Bool(m.AutoWheelCal)
		fields = append(fields, field)
	}
	if m.AutoPowerZero < 2 {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Bool(m.AutoPowerZero)
		fields = append(fields, field)
	}
	if m.Id != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(m.Id)
		fields = append(fields, field)
	}
	if m.SpdEnabled < 2 {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Bool(m.SpdEnabled)
		fields = append(fields, field)
	}
	if m.CadEnabled < 2 {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = proto.Bool(m.CadEnabled)
		fields = append(fields, field)
	}
	if m.SpdcadEnabled < 2 {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = proto.Bool(m.SpdcadEnabled)
		fields = append(fields, field)
	}
	if m.PowerEnabled < 2 {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = proto.Bool(m.PowerEnabled)
		fields = append(fields, field)
	}
	if m.CrankLength != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.Uint8(m.CrankLength)
		fields = append(fields, field)
	}
	if m.Enabled < 2 {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if m.BikeSpdAntIdTransType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint8(m.BikeSpdAntIdTransType)
		fields = append(fields, field)
	}
	if m.BikeCadAntIdTransType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint8(m.BikeCadAntIdTransType)
		fields = append(fields, field)
	}
	if m.BikeSpdcadAntIdTransType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Uint8(m.BikeSpdcadAntIdTransType)
		fields = append(fields, field)
	}
	if m.BikePowerAntIdTransType != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = proto.Uint8(m.BikePowerAntIdTransType)
		fields = append(fields, field)
	}
	if m.OdometerRollover != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 37)
		field.Value = proto.Uint8(m.OdometerRollover)
		fields = append(fields, field)
	}
	if m.FrontGearNum != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 38)
		field.Value = proto.Uint8(m.FrontGearNum)
		fields = append(fields, field)
	}
	if m.FrontGear != nil {
		field := fac.CreateField(mesg.Num, 39)
		field.Value = proto.SliceUint8(m.FrontGear)
		fields = append(fields, field)
	}
	if m.RearGearNum != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 40)
		field.Value = proto.Uint8(m.RearGearNum)
		fields = append(fields, field)
	}
	if m.RearGear != nil {
		field := fac.CreateField(mesg.Num, 41)
		field.Value = proto.SliceUint8(m.RearGear)
		fields = append(fields, field)
	}
	if m.ShimanoDi2Enabled < 2 {
		field := fac.CreateField(mesg.Num, 44)
		field.Value = proto.Bool(m.ShimanoDi2Enabled)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// OdometerScaled return Odometer in its scaled value.
// If Odometer value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *BikeProfile) OdometerScaled() float64 {
	if m.Odometer == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Odometer)/100 - 0
}

// CustomWheelsizeScaled return CustomWheelsize in its scaled value.
// If CustomWheelsize value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m
func (m *BikeProfile) CustomWheelsizeScaled() float64 {
	if m.CustomWheelsize == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.CustomWheelsize)/1000 - 0
}

// AutoWheelsizeScaled return AutoWheelsize in its scaled value.
// If AutoWheelsize value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m
func (m *BikeProfile) AutoWheelsizeScaled() float64 {
	if m.AutoWheelsize == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.AutoWheelsize)/1000 - 0
}

// BikeWeightScaled return BikeWeight in its scaled value.
// If BikeWeight value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: kg
func (m *BikeProfile) BikeWeightScaled() float64 {
	if m.BikeWeight == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BikeWeight)/10 - 0
}

// PowerCalFactorScaled return PowerCalFactor in its scaled value.
// If PowerCalFactor value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: %
func (m *BikeProfile) PowerCalFactorScaled() float64 {
	if m.PowerCalFactor == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.PowerCalFactor)/10 - 0
}

// CrankLengthScaled return CrankLength in its scaled value.
// If CrankLength value is invalid, float64 invalid value will be returned.
//
// Scale: 2; Offset: -110; Units: mm
func (m *BikeProfile) CrankLengthScaled() float64 {
	if m.CrankLength == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.CrankLength)/2 - -110
}

// SetMessageIndex sets MessageIndex value.
func (m *BikeProfile) SetMessageIndex(v typedef.MessageIndex) *BikeProfile {
	m.MessageIndex = v
	return m
}

// SetName sets Name value.
func (m *BikeProfile) SetName(v string) *BikeProfile {
	m.Name = v
	return m
}

// SetSport sets Sport value.
func (m *BikeProfile) SetSport(v typedef.Sport) *BikeProfile {
	m.Sport = v
	return m
}

// SetSubSport sets SubSport value.
func (m *BikeProfile) SetSubSport(v typedef.SubSport) *BikeProfile {
	m.SubSport = v
	return m
}

// SetOdometer sets Odometer value.
//
// Scale: 100; Units: m
func (m *BikeProfile) SetOdometer(v uint32) *BikeProfile {
	m.Odometer = v
	return m
}

// SetOdometerScaled is similar to SetOdometer except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *BikeProfile) SetOdometerScaled(v float64) *BikeProfile {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Odometer = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Odometer = uint32(unscaled)
	return m
}

// SetBikeSpdAntId sets BikeSpdAntId value.
//
// Base: uint16z
func (m *BikeProfile) SetBikeSpdAntId(v uint16) *BikeProfile {
	m.BikeSpdAntId = v
	return m
}

// SetBikeCadAntId sets BikeCadAntId value.
//
// Base: uint16z
func (m *BikeProfile) SetBikeCadAntId(v uint16) *BikeProfile {
	m.BikeCadAntId = v
	return m
}

// SetBikeSpdcadAntId sets BikeSpdcadAntId value.
//
// Base: uint16z
func (m *BikeProfile) SetBikeSpdcadAntId(v uint16) *BikeProfile {
	m.BikeSpdcadAntId = v
	return m
}

// SetBikePowerAntId sets BikePowerAntId value.
//
// Base: uint16z
func (m *BikeProfile) SetBikePowerAntId(v uint16) *BikeProfile {
	m.BikePowerAntId = v
	return m
}

// SetCustomWheelsize sets CustomWheelsize value.
//
// Scale: 1000; Units: m
func (m *BikeProfile) SetCustomWheelsize(v uint16) *BikeProfile {
	m.CustomWheelsize = v
	return m
}

// SetCustomWheelsizeScaled is similar to SetCustomWheelsize except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m
func (m *BikeProfile) SetCustomWheelsizeScaled(v float64) *BikeProfile {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.CustomWheelsize = uint16(basetype.Uint16Invalid)
		return m
	}
	m.CustomWheelsize = uint16(unscaled)
	return m
}

// SetAutoWheelsize sets AutoWheelsize value.
//
// Scale: 1000; Units: m
func (m *BikeProfile) SetAutoWheelsize(v uint16) *BikeProfile {
	m.AutoWheelsize = v
	return m
}

// SetAutoWheelsizeScaled is similar to SetAutoWheelsize except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m
func (m *BikeProfile) SetAutoWheelsizeScaled(v float64) *BikeProfile {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.AutoWheelsize = uint16(basetype.Uint16Invalid)
		return m
	}
	m.AutoWheelsize = uint16(unscaled)
	return m
}

// SetBikeWeight sets BikeWeight value.
//
// Scale: 10; Units: kg
func (m *BikeProfile) SetBikeWeight(v uint16) *BikeProfile {
	m.BikeWeight = v
	return m
}

// SetBikeWeightScaled is similar to SetBikeWeight except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: kg
func (m *BikeProfile) SetBikeWeightScaled(v float64) *BikeProfile {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BikeWeight = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BikeWeight = uint16(unscaled)
	return m
}

// SetPowerCalFactor sets PowerCalFactor value.
//
// Scale: 10; Units: %
func (m *BikeProfile) SetPowerCalFactor(v uint16) *BikeProfile {
	m.PowerCalFactor = v
	return m
}

// SetPowerCalFactorScaled is similar to SetPowerCalFactor except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: %
func (m *BikeProfile) SetPowerCalFactorScaled(v float64) *BikeProfile {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.PowerCalFactor = uint16(basetype.Uint16Invalid)
		return m
	}
	m.PowerCalFactor = uint16(unscaled)
	return m
}

// SetAutoWheelCal sets AutoWheelCal value.
func (m *BikeProfile) SetAutoWheelCal(v typedef.Bool) *BikeProfile {
	m.AutoWheelCal = v
	return m
}

// SetAutoPowerZero sets AutoPowerZero value.
func (m *BikeProfile) SetAutoPowerZero(v typedef.Bool) *BikeProfile {
	m.AutoPowerZero = v
	return m
}

// SetId sets Id value.
func (m *BikeProfile) SetId(v uint8) *BikeProfile {
	m.Id = v
	return m
}

// SetSpdEnabled sets SpdEnabled value.
func (m *BikeProfile) SetSpdEnabled(v typedef.Bool) *BikeProfile {
	m.SpdEnabled = v
	return m
}

// SetCadEnabled sets CadEnabled value.
func (m *BikeProfile) SetCadEnabled(v typedef.Bool) *BikeProfile {
	m.CadEnabled = v
	return m
}

// SetSpdcadEnabled sets SpdcadEnabled value.
func (m *BikeProfile) SetSpdcadEnabled(v typedef.Bool) *BikeProfile {
	m.SpdcadEnabled = v
	return m
}

// SetPowerEnabled sets PowerEnabled value.
func (m *BikeProfile) SetPowerEnabled(v typedef.Bool) *BikeProfile {
	m.PowerEnabled = v
	return m
}

// SetCrankLength sets CrankLength value.
//
// Scale: 2; Offset: -110; Units: mm
func (m *BikeProfile) SetCrankLength(v uint8) *BikeProfile {
	m.CrankLength = v
	return m
}

// SetCrankLengthScaled is similar to SetCrankLength except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 2; Offset: -110; Units: mm
func (m *BikeProfile) SetCrankLengthScaled(v float64) *BikeProfile {
	unscaled := (v + -110) * 2
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.CrankLength = uint8(basetype.Uint8Invalid)
		return m
	}
	m.CrankLength = uint8(unscaled)
	return m
}

// SetEnabled sets Enabled value.
func (m *BikeProfile) SetEnabled(v typedef.Bool) *BikeProfile {
	m.Enabled = v
	return m
}

// SetBikeSpdAntIdTransType sets BikeSpdAntIdTransType value.
//
// Base: uint8z
func (m *BikeProfile) SetBikeSpdAntIdTransType(v uint8) *BikeProfile {
	m.BikeSpdAntIdTransType = v
	return m
}

// SetBikeCadAntIdTransType sets BikeCadAntIdTransType value.
//
// Base: uint8z
func (m *BikeProfile) SetBikeCadAntIdTransType(v uint8) *BikeProfile {
	m.BikeCadAntIdTransType = v
	return m
}

// SetBikeSpdcadAntIdTransType sets BikeSpdcadAntIdTransType value.
//
// Base: uint8z
func (m *BikeProfile) SetBikeSpdcadAntIdTransType(v uint8) *BikeProfile {
	m.BikeSpdcadAntIdTransType = v
	return m
}

// SetBikePowerAntIdTransType sets BikePowerAntIdTransType value.
//
// Base: uint8z
func (m *BikeProfile) SetBikePowerAntIdTransType(v uint8) *BikeProfile {
	m.BikePowerAntIdTransType = v
	return m
}

// SetOdometerRollover sets OdometerRollover value.
//
// Rollover counter that can be used to extend the odometer
func (m *BikeProfile) SetOdometerRollover(v uint8) *BikeProfile {
	m.OdometerRollover = v
	return m
}

// SetFrontGearNum sets FrontGearNum value.
//
// Base: uint8z; Number of front gears
func (m *BikeProfile) SetFrontGearNum(v uint8) *BikeProfile {
	m.FrontGearNum = v
	return m
}

// SetFrontGear sets FrontGear value.
//
// Base: uint8z; Array: [N]; Number of teeth on each gear 0 is innermost
func (m *BikeProfile) SetFrontGear(v []uint8) *BikeProfile {
	m.FrontGear = v
	return m
}

// SetRearGearNum sets RearGearNum value.
//
// Base: uint8z; Number of rear gears
func (m *BikeProfile) SetRearGearNum(v uint8) *BikeProfile {
	m.RearGearNum = v
	return m
}

// SetRearGear sets RearGear value.
//
// Base: uint8z; Array: [N]; Number of teeth on each gear 0 is innermost
func (m *BikeProfile) SetRearGear(v []uint8) *BikeProfile {
	m.RearGear = v
	return m
}

// SetShimanoDi2Enabled sets ShimanoDi2Enabled value.
func (m *BikeProfile) SetShimanoDi2Enabled(v typedef.Bool) *BikeProfile {
	m.ShimanoDi2Enabled = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *BikeProfile) SetUnknownFields(unknownFields ...proto.Field) *BikeProfile {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *BikeProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *BikeProfile {
	m.DeveloperFields = developerFields
	return m
}