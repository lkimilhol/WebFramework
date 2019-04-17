package message

type requestMessage struct {
	apiName string
	params map[string]interface{}
}

//func MakeMessageForApiCall(apiName string, params interface{}) requestMessage {
//	dataMap := paramsJsonToMap(params)
//	message := requestMessage{apiName:apiName, params:dataMap}
//	return message
//}
//
//func paramsJsonToMap(params interface{}) map[string]interface{} {
//	dataMap := make(map[string]interface{})
//	err := json.Unmarshal([]byte(params), &dataMap)
//
//	if err != nil {
//		panic(err)
//	}
//
//	return dataMap
//}