package testing

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TokenName struct {
	Name string `json:"name"`
}

func TestCreateNewToken(url, name, method string) []byte {
	client := http.Client{}
	tn := TokenName{
		Name: name,
	}
	bodyReq, err := json.Marshal(tn)
	req, errRequest := http.NewRequest(method, url, bytes.NewBuffer(bodyReq))
	if errRequest != nil {
		log.Println(errRequest)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	return body
}

type TokenStruct struct {
	Token string `json:"token"`
}

func TestTokenCheckInfo(url, token, method string) []byte {
	client := http.Client{}
	tn := TokenStruct{
		Token: token,
	}
	bodyReq, err := json.Marshal(tn)
	req, errRequest := http.NewRequest(method, url, bytes.NewBuffer(bodyReq))
	if errRequest != nil {
		log.Println(errRequest)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	return body
}
