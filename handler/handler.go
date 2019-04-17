package handler
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func makeInstance(name string) interface{} {
//	v := reflect.New(typeRegistry[name]).Elem()
//	return v.Interface()
//}
//
//
//func (wasServer *WasServer) handlerMapping() {
//	wasServer.HandlerGroup = make(map[string]func(gatewayMessage &message.GatewayMessage))
//
//	wasServer.HandlerGroup["Login"] = handler.LoginHandler
//}
//
////TODO 핸들러를 여기에 등록 할 수 있도록?
//func MakeInstance(api interface{}) {
//	//todo handler의 List 구조체 안에 있는 정보들을 모두 map에 담도록?
//
//	target := reflect.ValueOf(api)
//	fmt.Println(reflect.TypeOf(target))
//	elements := target.Elem()
//
//	for i := 0; i < elements.NumField(); i++ {
//		mValue := elements.Field(i)
//		mType := elements.Type().Field(i)
//		tag := mType.Tag
//
//		fmt.Printf("%10s %10s ==> %10v, json: %10s\n",
//			mType.Name,         // 이름
//			mType.Type,         // 타입
//			mValue.Interface(), // 값
//			tag.Get("json"))    // json 태그
//	}
//}
//
//
