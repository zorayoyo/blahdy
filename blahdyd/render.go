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

func RenderBlah(id string) []byte {
	blahTmp := T_BLAH{}
	blah := BlahdyDB.GetBlah(id)
	user := BlahdyDB.GetUser(blah.AuthorId)
	if user == nil {
		return nil
	}
	blahTmp.Build(blah, user)
	return RenderJson(blahTmp)
}


func RenderTimeline(id string) []byte {
	msgTmps := make([]*T_MESSAGE, 0)
	msgs := BlahdyDB.GetTimeline(id)
	for _, msg := range msgs {
		user := BlahdyDB.GetUser(msg.AuthorId)
		if user == nil {
			continue
		}
		mt := new(T_MESSAGE)
		mt.Build(msg, user)
		msgTmps = append(msgTmps, mt)
	}
	return RenderJson(msgTmps)
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

