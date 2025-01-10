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
		v := map[string]string{}
		if r.Ping.Content != "" {
			v["ping"] = r.Ping.Content
		}
		if r.DNS.Content != "" {
			v["dns"] = r.DNS.Content
		}
		if r.SystemDNS.Content != "" {
			v["system_dns"] = r.SystemDNS.Content
		}
		if r.TLS.Content != "" {
			v["tls"] = r.TLS.Content
		}
		if r.HTTPS.Content != "" {
			v["https"] = r.HTTPS.Content
		}

		byt, _ := json.MarshalIndent(v, "", "  ")
		output += string(byt)
	}

	return output
}
