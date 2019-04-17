package api

import "WebFramework/handler"

func GetApiList() map[string]interface{} {
	apiList := map[string]interface{}{
		"login": handler.Login{},
	}
	return apiList
}




















//type List struct {
//	login handler.Login
//}
//
//func (list *List) Init() map[string]interface{} {
//	target := reflect.ValueOf(list)
//	elements := target.Elem()
//
//	apiList := make(map[string]interface{})
//
//	for i := 0; i < elements.NumField(); i++ {
//		//mValue := elements.Field(i)
//		//mType := reflect.TypeOf(mValue)
//		mKey := elements.Type().Field(i)
//		//ptr := reflect.New(elements.Type())
//		//tmp := reflect.New(mType)
//		a := reflect.New(reflect.TypeOf(mKey))
//
//		apiList[mKey.Name] = a
//	}
//
//	return apiList
//}
