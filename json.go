package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Members interface{}
type data struct {
	Members `json:"members"`
}

func main() {

	resp, err := http.Get("https://gist.githubusercontent.com/DQIJAO/e14c64ea610688e70228a9fb8c649b2c/raw/6cccd444c1ef65411aa3662b112634996b837414/bnk48.json")

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	var data data

	err = json.Unmarshal(respByte, &data)
	if err != nil {
		fmt.Println(err)

		return
	}

	log.Println(data)

}
