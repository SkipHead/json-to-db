package controllers

import (
	"encoding/json"
	"io/ioutil"
	"json-to-postgres/internal/app/parser/models"
	"log"
)

func Parse(b []byte) models.JsonStruct {
	j := models.JsonStruct{}
	err := json.Unmarshal(b, &j)
	if err != nil {
		log.Println(err)
	}

	return j
}

func OpenFileJson(path string) []byte {
	openJson, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return openJson
}
