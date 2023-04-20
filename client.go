package capGo

func Client(ApiKey string) *capGoStruct{
	capGo := &capGoStruct{}
	capGo.ApiKey = ApiKey
   
	return capGo
}