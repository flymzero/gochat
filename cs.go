package mzgochat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

const (
	configPath = "./config.yaml"
)

const (
	ClientPlatform = iota
	ServerPlatform
)

const (
	MessageInfoType = iota
	ServerInfoType
	ClientInfoType
)

type CSData struct {
	NickName string `json:nickname`
	Info     string `json:info`
	InfoType int    `json:infotype`
	Platform int    `json:playform`
}

func (c CSData) ToJson() []byte {
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Print(err)
	}
	return bytes
}

func UnmarshalYaml() {
	cmd := exec.Command("pwd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	if data, err := ioutil.ReadFile(configPath); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(data)
	}

}
