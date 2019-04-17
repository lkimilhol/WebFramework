package main

import (
	"WebFramework/api"
	"WebFramework/config"
	_ "WebFramework/handler"
	"WebFramework/message"
	_ "WebFramework/util"
	"fmt"
	"net/http"
	"reflect"
)

var apiList map[string]interface{}

func init() {
	apiList = api.GetApiList()
}

func main() {
	addr := serverInit()

	http.Handle("/", new(testHandler))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	packet := message.JsonDecoder(req)

	//클라이언트 버전 등등을 체크 하는 로직도 필요

	object := apiList[packet.ApiName]


	//TODO JSON Parsing 후에 API Name을 따라서 호출하는 부분 해결해야함

	//requestParam := message.MakeMessageForApiCall(packet.ApiName, packet.Params)
	//fmt.Println(requestParam)

	//a := apiList[packet.ApiName]
	//reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})


	//new(a)



	//fmt.Println(a)
	/**
	임시로 패킷의 내용을 보기 위해 남겨둠
	*/
	fmt.Println(packet.ClientVersion)
	fmt.Println(packet.Params)
	fmt.Println(packet.Timestamp)

	fmt.Println(reflect.TypeOf(packet.Params))

	//TODO 분기 처리 어떻게 할지 생각해야함
}

func serverInit() string {
	configInfo := config.ReadConfigFile()
	addr := configInfo.IP + ":" + configInfo.Port
	return addr
}
