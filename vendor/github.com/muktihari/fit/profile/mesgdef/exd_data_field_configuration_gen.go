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
)

// ExdDataFieldConfiguration is a ExdDataFieldConfiguration message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type ExdDataFieldConfiguration struct {
	Title        [32]string // Array: [32]
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptCount uint8
	DisplayType  typedef.ExdDisplayType

	state [1]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewExdDataFieldConfiguration creates new ExdDataFieldConfiguration struct based on given mesg.
// If mesg is nil, it will return ExdDataFieldConfiguration with all fields being set to its corresponding invalid value.
func NewExdDataFieldConfiguration(mesg *proto.Message) *ExdDataFieldConfiguration {
	vals := [6]proto.Value{}

	var state [1]uint8
	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 5 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 4 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &ExdDataFieldConfiguration{
		ScreenIndex:  vals[0].Uint8(),
		ConceptField: vals[1].Uint8(),
		FieldId:      vals[2].Uint8(),
		ConceptCount: vals[3].Uint8(),
		DisplayType:  typedef.ExdDisplayType(vals[4].Uint8()),
		Title: func() (arr [32]string) {
			arr = [32]string{
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
				basetype.StringInvalid,
			}
			copy(arr[:], vals[5].SliceString())
			return arr
		}(),

		state: state,

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts ExdDataFieldConfiguration into proto.Message. If options is nil, default options will be used.
func (m *ExdDataFieldConfiguration) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumExdDataFieldConfiguration}

	if m.ScreenIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.ScreenIndex)
		fields = append(fields, field)
	}
	if m.ConceptField != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.ConceptField)
		fields = append(fields, field)
	}
	if m.FieldId != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(2); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 2)
			field.Value = proto.Uint8(m.FieldId)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.ConceptCount != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(3); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 3)
			field.Value = proto.Uint8(m.ConceptCount)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.DisplayType != typedef.ExdDisplayTypeInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.DisplayType))
		fields = append(fields, field)
	}
	if m.Title != [32]string{
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
		basetype.StringInvalid,
	} {
		field := fac.CreateField(mesg.Num, 5)
		copied := m.Title
		field.Value = proto.SliceString(copied[:])
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

// SetScreenIndex sets ScreenIndex value.
func (m *ExdDataFieldConfiguration) SetScreenIndex(v uint8) *ExdDataFieldConfiguration {
	m.ScreenIndex = v
	return m
}

// SetConceptField sets ConceptField value.
func (m *ExdDataFieldConfiguration) SetConceptField(v byte) *ExdDataFieldConfiguration {
	m.ConceptField = v
	return m
}

// SetFieldId sets FieldId value.
func (m *ExdDataFieldConfiguration) SetFieldId(v uint8) *ExdDataFieldConfiguration {
	m.FieldId = v
	return m
}

// SetConceptCount sets ConceptCount value.
func (m *ExdDataFieldConfiguration) SetConceptCount(v uint8) *ExdDataFieldConfiguration {
	m.ConceptCount = v
	return m
}

// SetDisplayType sets DisplayType value.
func (m *ExdDataFieldConfiguration) SetDisplayType(v typedef.ExdDisplayType) *ExdDataFieldConfiguration {
	m.DisplayType = v
	return m
}

// SetTitle sets Title value.
//
// Array: [32]
func (m *ExdDataFieldConfiguration) SetTitle(v [32]string) *ExdDataFieldConfiguration {
	m.Title = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *ExdDataFieldConfiguration) SetUnknownFields(unknownFields ...proto.Field) *ExdDataFieldConfiguration {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *ExdDataFieldConfiguration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExdDataFieldConfiguration {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 2, 3.
func (m *ExdDataFieldConfiguration) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 2, 3:
	default:
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	m.state[pos] &^= bit
	if flag {
		m.state[pos] |= bit
	}
	return true
}

// IsExpandedField checks whether given fieldNum is a field generated through
// a component expansion. Eligible for field number: 2, 3.
func (m *ExdDataFieldConfiguration) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 4 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}