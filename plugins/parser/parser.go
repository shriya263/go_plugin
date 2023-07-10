package main

type speaker struct {
	Message string
	Number  int
}

func (s *speaker) GetDetails() interface{} {
	return s
}

func (s *speaker) Speak() string {
	return "hello"
}

type parser struct {
}

func (p *parser) GetLatitude() map[string]interface{} {
	latDetails := make(map[string]interface{})
	latDetails["StartIndex"] = 1
	latDetails["EndIndex"] = 2
	return latDetails
}

// Exported
var Parser parser
var Speaker speaker
var SpeakerName = "Alice"

func init() {
	Speaker.Message = "hello "
	Speaker.Number = 12
}
