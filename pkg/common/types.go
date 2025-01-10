package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

type Args []string

type Address string

type ResultPart struct {
	Content string
	Error   error
}

func (o ResultPart) String() string {
	if o.Error != nil {
		return Red(o.Error.Error())
	}

	return o.Content
}

func (a Address) String() string {
	return string(a)
}

type Result struct {
	Ping      ResultPart `json:"ping"`
	DNS       ResultPart `json:"dns"`
	SystemDNS ResultPart `json:"system_dns"`
	TLS       ResultPart `json:"tls"`
	HTTPS     ResultPart `json:"https"`

	Mu sync.Mutex `json:"-"`
}

func (r Result) Output(outputType string) string {
	var output string

	switch outputType {
	case "plaintext":
		v := reflect.ValueOf(r)
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).Name == "Mu" {
				continue
			}

			output += fmt.Sprintf("%s: %v\n", White(v.Type().Field(i).Name), Green(v.Field(i).Interface()))
		}
	case "json":
		// SAFETY: we are sure that this never fails
		byt, _ := json.MarshalIndent(r, "", "  ")
		output += string(byt)
	}

	return output
}
