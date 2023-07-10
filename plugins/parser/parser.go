package main

import (
	"encoding/json"
	"fmt"
)

type speaker struct {
	Message string
	Number  int
}

type Longitude struct {
	StartIndex int
	EndIndex   int
	DataType   string
}

func (s *speaker) GetDetails() interface{} {
	return s
}

func (s *speaker) Speak() string {
	return "hello"
}

type parser struct {
}

// Approach 1: creating a map for fields
func (p *parser) GetLatitude() map[string]interface{} {
	latDetails := make(map[string]interface{})
	latDetails["StartIndex"] = 1
	latDetails["EndIndex"] = 2
	return latDetails
}

// Approach 2: using json marshal
func (p *parser) GetLongitude() string {
	longitudeBytes, err := json.Marshal(Longitude{StartIndex: 2, EndIndex: 3, DataType: "string"})
	if err != nil {
		fmt.Println("unable to parser ", err)
		return ""
	}
	return string(longitudeBytes)
}

// Exported
var Parser parser
var Speaker speaker
var SpeakerName = "Alice"

func init() {
	Speaker.Message = "hello "
	Speaker.Number = 12
}
