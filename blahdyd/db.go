package main

import (
	"github.com/shellex/tattoo/webapp"
	"encoding/json"
	"strconv"
	"fmt"
)

type BlahdyStorage struct {
	BlahDB  webapp.FileStorage
	UserDB	webapp.FileStorage
	ActionDB	webapp.FileStorage
	VarDB   webapp.FileStorage
}

var BlahdyDB * BlahdyStorage = nil

func init() {
	BlahdyDB = new(BlahdyStorage)
}

func (db * BlahdyStorage) Load(app *webapp.App) {
	app.Log("DB", "Init DB: Blah DB")
	db.BlahDB.Init("storage/blah/", webapp.FILE_STORAGE_MODE_MULIPLE)
	app.Log("DB", "Init DB: User DB")
	db.UserDB.Init("storage/user/", webapp.FILE_STORAGE_MODE_MULIPLE)
	app.Log("DB", "Init DB: Action DB")
	db.ActionDB.Init("storage/action/", webapp.FILE_STORAGE_MODE_MULIPLE)
	app.Log("DB", "Init DB: Vars DB")
	db.VarDB.Init("storage/var/", webapp.FILE_STORAGE_MODE_MULIPLE)
	// samples
	for i := 0; i < 3; i += 1 {
		blah := Blah{
			strconv.Itoa(i), "ABC #" + strconv.Itoa(i), "1", 1, 1,
		}
		BlahdyDB.CreateBlah(&blah)
	}
}

func (db * BlahdyStorage) GetFreeIdByName(name string) string {
	// read the current blah id from var db
	// increase it.
	// write back and return.
	// i may meet a race problem?
	var current uint64
	var err error
	currentStr, err := db.VarDB.GetString(name)
	if err != nil {
		current = 1
	} else {
		current, err = strconv.ParseUint(currentStr, 10, 64)
		if err != nil {
			return "0"
		}
		current += 1
	}
	currentStr = strconv.FormatUint(current, 10)
	fmt.Printf("%v\n",currentStr)
	db.VarDB.SetString(name, currentStr)
	db.VarDB.SaveIndex()
	return currentStr;
}

func (db * BlahdyStorage) GetFreeBlahId() string {
	return db.GetFreeIdByName("freeBlahId")
}

func (db * BlahdyStorage) GetFreeUserId() string {
	return db.GetFreeIdByName("freeUserId")
}

func (db * BlahdyStorage) CreateBlah(blah * Blah) ([]byte, error) {
	// get a free id
	// create a item in BlahDB with this id as key
	// return the content of the new item
	freeId := db.GetFreeBlahId()
	blah.Id = freeId
	blahJson, err := json.Marshal(blah)
	if err != nil {
		return nil, err
	}
	db.BlahDB.Set(freeId, blahJson)
	db.BlahDB.SaveIndex()
	return blahJson, nil
}

func (db * BlahdyStorage) DestroyBlah(id string) ([]byte, error) {
	// check existence of the specified id
	// and remove it from db
	blah, err := db.BlahDB.Get(id)
	if blah != nil {
		db.BlahDB.Delete(id)
		db.BlahDB.SaveIndex()
	}
	return blah, err
}

func (db * BlahdyStorage) GetBlah(id string) ([]byte, error) {
	// @TODO a custom error should be defined.
	if ! db.BlahDB.Has(id) {
		return nil, nil
	}
	blah, _ := db.BlahDB.Get(id)
	return blah, nil
}

func (db * BlahdyStorage) GetBlahJson(id string) (interface{}, error) {
	if ! db.BlahDB.Has(id) {
		return nil, nil
	}
	blah, _ := db.BlahDB.GetJSON(id)
	return blah, nil
}

func (db * BlahdyStorage) GetBlahs() []*Blah {
	blahs := make([]*Blah, 0)
	for id, _ := range db.BlahDB.Index {
		if id == "*" {
			continue
		}
		blahJson, _ := db.GetBlahJson(id)
		blah := new(Blah)
		blah.BuildFromJson(blahJson)
		blahs = append(blahs, blah)
	}
	return blahs
}


