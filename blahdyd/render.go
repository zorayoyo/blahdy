package main

import (
    "encoding/json"
	"os"
    "io/ioutil"
	"github.com/shellex/tattoo/webapp"
)

var JsonSample map[string][]byte;

func LoadSamples() {
	samples := []string {
		"blah/all",
		"blah/actions",
		"blah/members",
		"blah/show",
	}
	JsonSample = make(map[string][]byte)
	for _, name := range samples {
		filename := "json_sample/" + name + ".json"
		if _, err := os.Stat(filename); err == nil {
			data, _ := ioutil.ReadFile(filename)
			JsonSample[name] = data;
		}
	}
}

func RenderJsonSample(ctx * webapp.Context, tplName string) []byte {
    if value, ok := JsonSample[tplName]; ok {
        return value
    }
	return nil
}

func RenderJson(ctx * webapp.Context, value interface{}) []byte {
	blahJson, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	return blahJson
}

