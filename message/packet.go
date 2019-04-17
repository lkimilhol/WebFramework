package message

import (
	"encoding/json"
	"net/http"
)

type Packet struct {
	ClientVersion string `json:"client_version"`
	ApiName       string `json:"api_name"`
	Params        interface{} `json:"params"`
	Timestamp int64 `json:"timestamp"`
}

//TODO 포인터 변경 후 nil 체크 되도록 수정


func JsonDecoder(request *http.Request) Packet {
	decoder := json.NewDecoder(request.Body)

	var packet Packet
	err := decoder.Decode(&packet)

	if err != nil {
		panic(err)
	}
	return packet
}
