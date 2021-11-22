package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

type Output string

type Args []string

type Address string

func (o Output) String() string {
	return string(o)
}

func (a Address) String() string {
	return string(a)
}

type Result struct {
	Ping      string `json:"ping"`
	DNS       string `json:"dns"`
	SystemDNS string `json:"system_dns"`
	TLS       string `json:"tls"`
	HTTPS     string `json:"https"`

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
