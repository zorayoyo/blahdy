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

func RenderAllBlahs() []byte {
	blahsTmps := make([]*T_BLAH, 0)
	blahs := BlahdyDB.GetBlahs()
	for _, blah := range blahs {
		user := BlahdyDB.GetUser(blah.AuthorId)
		if user == nil {
			continue
		}
		bt := new(T_BLAH)
		bt.Build(blah, user)
		blahsTmps = append(blahsTmps, bt)
	}
	return RenderJson(blahsTmps)
}

func RenderJsonSample(ctx * webapp.Context, tplName string) []byte {
    if value, ok := JsonSample[tplName]; ok {
        return value
    }
	return nil
}

func RenderJson(value interface{}) []byte {
	blahJson, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	return blahJson
}

