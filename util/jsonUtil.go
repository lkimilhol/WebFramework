package util

// request 요청을 json 으로 decode 해주는 부분
// 값이 없을 경우 여기서 각 자료형에 맞는 default value 로 셋팅이 됨
//func JsonDecoder(request *http.Request) *message.Packet {
//	packet := new (message.Packet)
//	err := json.NewDecoder(request).Decode(&packet)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	//TODO 디코더는 유틸일 뿐이고 체크는 전부 각 object?? 들에서 해줘야 할 듯
//
//	return packet
//}