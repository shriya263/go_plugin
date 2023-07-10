package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"plugin"
	"strings"
)

type Speaker interface {
	Speak() string
	GetDetails() interface{}
}

type Parser interface {
	GetLatitude() map[string]interface{}
	GetLongitude() string
}

type longitude struct {
	StartIndex int
	EndIndex   int
	DataType   string
}

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run(args []string) error {
	pulginName := "english"
	if len(args) == 2 {
		pulginName = args[1]
	}
	var mod string
	switch pulginName {
	case "english":
		mod = "./plugins/eng/eng.so"
	case "vietnamese":
		mod = "plugins/vie/vie.so"
	case "parser":
		mod = "plugins/parser/parser.so"
	default:
		return errors.New("this speakerName is not supported")
	}
	// To open the plugin which is mentioned in path (i.e mod )
	plugin, err := plugin.Open(mod)
	if err != nil {
		return err
	}

	speaker, err := lookUpSymbol[Speaker](plugin, "Speaker")
	if err != nil {
		return err
	}

	speakerName, err := lookUpSymbol[string](plugin, "SpeakerName")
	if err != nil {
		return err
	}
	fmt.Printf("%s says \"%s\" in %s \n", *speakerName, (*speaker).Speak(), strings.Title(pulginName))

	if pulginName == "parser" {
		parser, err := lookUpSymbol[Parser](plugin, "Parser")
		if err != nil {
			return err
		}
		fmt.Println("--------")
		// Approach 1: Through map[string]interface{}
		fmt.Println("Through map")
		latitudeDetails := (*parser).GetLatitude()
		fmt.Println("latitude ", latitudeDetails)

		// Approach 2 : Through json
		fmt.Println("Through json")
		longitudeDetails := (*parser).GetLongitude()

		var longitudeData longitude
		_ = json.Unmarshal([]byte(longitudeDetails), &longitudeData)
		fmt.Println("longitude details ", longitudeData.StartIndex, " ", longitudeData.EndIndex)
	}

	return nil
}

func lookUpSymbol[M any](plugin *plugin.Plugin, symbolName string) (*M, error) {
	// Lookup Func : will go the specified plugin and gather the Varible and searches in plugines exported variable
	symbol, err := plugin.Lookup(symbolName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("type symbol: %T\n ", symbol)
	switch symbol.(type) {
	case *M:
		// for primitives
		return symbol.(*M), nil
	case M:
		// for non-primitive
		result := symbol.(M)
		return &result, nil
	default:
		return nil, errors.New(fmt.Sprintf("unexpected type from module symbol: %T", symbol))
	}
}
