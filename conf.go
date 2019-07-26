package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type configuration struct {
	BotToken string `json:"bot_token"`
	Prefix   string `json:"prefix"`
}

var conf *configuration

func loadConfig() (err error) {
	cf, err := os.Open("conf.json")
	if err != nil {
		return
	}
	defer cf.Close()

	data, _ := ioutil.ReadAll(cf)
	conf = &configuration{}
	json.Unmarshal(data, &conf)

	return
}
