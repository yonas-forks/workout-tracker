// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// CoursePoint is a CoursePoint message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type CoursePoint struct {
	Timestamp    time.Time
	Name         string
	PositionLat  int32  // Units: semicircles
	PositionLong int32  // Units: semicircles
	Distance     uint32 // Scale: 100; Units: m
	MessageIndex typedef.MessageIndex
	Type         typedef.CoursePoint
	Favorite     typedef.Bool

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewCoursePoint creates new CoursePoint struct based on given mesg.
// If mesg is nil, it will return CoursePoint with all fields being set to its corresponding invalid value.
func NewCoursePoint(mesg *proto.Message) *CoursePoint {
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

	return &CoursePoint{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Timestamp:    datetime.ToTime(vals[1].Uint32()),
		PositionLat:  vals[2].Int32(),
		PositionLong: vals[3].Int32(),
		Distance:     vals[4].Uint32(),
		Type:         typedef.CoursePoint(vals[5].Uint8()),
		Name:         vals[6].String(),
		Favorite:     vals[8].Bool(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts CoursePoint into proto.Message. If options is nil, default options will be used.
func (m *CoursePoint) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumCoursePoint}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Int32(m.PositionLat)
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Int32(m.PositionLong)
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.Distance)
		fields = append(fields, field)
	}
	if m.Type != typedef.CoursePointInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}
	if m.Favorite < 2 {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Bool(m.Favorite)
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

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *CoursePoint) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// DistanceScaled return Distance in its scaled value.
// If Distance value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *CoursePoint) DistanceScaled() float64 {
	if m.Distance == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Distance)/100 - 0
}

// PositionLatDegrees returns PositionLat in degrees instead of semicircles.
// If PositionLat value is invalid, float64 invalid value will be returned.
func (m *CoursePoint) PositionLatDegrees() float64 {
	return semicircles.ToDegrees(m.PositionLat)
}

// PositionLongDegrees returns PositionLong in degrees instead of semicircles.
// If PositionLong value is invalid, float64 invalid value will be returned.
func (m *CoursePoint) PositionLongDegrees() float64 {
	return semicircles.ToDegrees(m.PositionLong)
}

// SetMessageIndex sets MessageIndex value.
func (m *CoursePoint) SetMessageIndex(v typedef.MessageIndex) *CoursePoint {
	m.MessageIndex = v
	return m
}

// SetTimestamp sets Timestamp value.
func (m *CoursePoint) SetTimestamp(v time.Time) *CoursePoint {
	m.Timestamp = v
	return m
}

// SetPositionLat sets PositionLat value.
//
// Units: semicircles
func (m *CoursePoint) SetPositionLat(v int32) *CoursePoint {
	m.PositionLat = v
	return m
}

// SetPositionLatDegrees is similar to SetPositionLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *CoursePoint) SetPositionLatDegrees(degrees float64) *CoursePoint {
	m.PositionLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetPositionLong sets PositionLong value.
//
// Units: semicircles
func (m *CoursePoint) SetPositionLong(v int32) *CoursePoint {
	m.PositionLong = v
	return m
}

// SetPositionLongDegrees is similar to SetPositionLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *CoursePoint) SetPositionLongDegrees(degrees float64) *CoursePoint {
	m.PositionLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetDistance sets Distance value.
//
// Scale: 100; Units: m
func (m *CoursePoint) SetDistance(v uint32) *CoursePoint {
	m.Distance = v
	return m
}

// SetDistanceScaled is similar to SetDistance except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *CoursePoint) SetDistanceScaled(v float64) *CoursePoint {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Distance = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Distance = uint32(unscaled)
	return m
}

// SetType sets Type value.
func (m *CoursePoint) SetType(v typedef.CoursePoint) *CoursePoint {
	m.Type = v
	return m
}

// SetName sets Name value.
func (m *CoursePoint) SetName(v string) *CoursePoint {
	m.Name = v
	return m
}

// SetFavorite sets Favorite value.
func (m *CoursePoint) SetFavorite(v typedef.Bool) *CoursePoint {
	m.Favorite = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *CoursePoint) SetUnknownFields(unknownFields ...proto.Field) *CoursePoint {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *CoursePoint) SetDeveloperFields(developerFields ...proto.DeveloperField) *CoursePoint {
	m.DeveloperFields = developerFields
	return m
}