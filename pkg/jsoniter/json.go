package jsoniter

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	Json jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary
)

func GetJsoniter() jsoniter.API {
	return Json
}
