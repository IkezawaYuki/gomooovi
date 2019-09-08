package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func init(){
	loadConfig()
}

// 設定ファイルの読み込み
func loadConfig(){
	file, err := os.Open("config.json")
	if err != nil{
		log.Fatalln(err)
	}
	decorder := json.NewDecoder(file)
	var config Configuration
	err = decorder.Decode(&config)
	if err != nil{
		log.Fatalln(err)
	}
}


