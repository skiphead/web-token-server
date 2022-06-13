package Server

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

//configStruct Struct for config json file
type configStruct struct {
	Port      string `json:"Port"`
	ExpiredAt int    `json:"expired_at"`
	TLS       bool   `json:"ssl"`
	ServerCrt string `json:"server_crt"`
	ServerKey string `json:"server_key"`
}

//Variables for config
var (
	configPath string
	conf       configStruct
)

// Init initialize config
func init() {
	flag.StringVar(&configPath, "config-path", "./config/server.json", "config path file")
	conf = OpenConf()
}

// OpenConf parsing json config file
func OpenConf() configStruct {

	flag.Parse()
	config := configStruct{}

	configFile, errOpenConfigJson := os.Open(configPath)
	if errOpenConfigJson != nil {
		log.Println("Error open config file", errOpenConfigJson)
	}

	log.Println("Successfully Opened configStruct", configPath)

	defer func(configFile *os.File) {
		errConfigFile := configFile.Close()
		if errConfigFile != nil {
			log.Println("Error close config file", errConfigFile)
		}
	}(configFile)

	byteConfig, errByteConfig := ioutil.ReadAll(configFile)
	if errByteConfig != nil {
		log.Println("Err read byte", errByteConfig)
	}

	err := json.Unmarshal(byteConfig, &config)
	if err != nil {
		log.Println("Error parse config", err)
	}

	return config
}
