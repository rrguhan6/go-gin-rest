package main

import (
	"bytes"
	"encoding/json"
	"regexp"
)

var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z]{1,})`)

type conventionalMarshaller struct {
	Value interface{}
}

func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)

	return converted, err
}

type DATA struct {
	Ev     string `json:"ev" binding:"required"`
	Et     string `json:"et" binding:"required"`
	ID     string `json:"id" binding:"required"`
	UID    string `json:"uid" binding:"required"`
	Mid    string `json:"mid" binding:"required"`
	T      string `json:"t" binding:"required"`
	P      string `json:"p" binding:"required"`
	L      string `json:"l" binding:"required"`
	Sc     string `json:"sc" binding:"required"`
	Atrk1  string `json:"atrk1" binding:"required"`
	Atrv1  string `json:"atrv1" binding:"required"`
	Atrt1  string `json:"atrt1" binding:"required"`
	Atrk2  string `json:"atrk2" binding:"required"`
	Atrv2  string `json:"atrv2" binding:"required"`
	Atrt2  string `json:"atrt2" binding:"required"`
	Uatrk1 string `json:"uatrk1" binding:"required"`
	Uatrv1 string `json:"uatrv1" binding:"required"`
	Uatrt1 string `json:"uatrt1" binding:"required"`
	Uatrk2 string `json:"uatrk2" binding:"required"`
	Uatrv2 string `json:"uatrv2" binding:"required"`
	Uatrt2 string `json:"uatrt2" binding:"required"`
	Uatrk3 string `json:"uatrk3" binding:"required"`
	Uatrv3 string `json:"uatrv3" binding:"required"`
	Uatrt3 string `json:"uatrt3" binding:"required"`
}

type ValueType struct {
	Value string
	Type  string
}

type Attributes struct {
	FormVarient ValueType
	Ref         ValueType
}

type Traits struct {
	Name  ValueType
	Email ValueType
	Age   ValueType
}

type NestedData struct {
	Event           string
	EventType       string
	AppID           string
	UserID          string
	MessageID       string
	PageTitle       string
	PageURL         string
	BrowserLanguage string
	ScreenSize      string
	Attributes      Attributes
	Traits          Traits
}
