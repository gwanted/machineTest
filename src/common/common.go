package common

import (
	"encoding/json"
	"net/http"
)

type R struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type RE struct {
	Code int64 `json:"code"`
	Msg  string`json:"msg"`
}


func ReturnFormat(w http.ResponseWriter, code int64, data interface{}, msg string) {
	res := R{Code:code, Data:data, Msg:msg}
	omg, _ := json.Marshal(res)
	w.Write(omg)
}

func ReturnEFormat(w http.ResponseWriter, code int64, msg string) {
	res := RE{Code:code, Msg:msg}
	omg, _ := json.Marshal(res)
	w.Write(omg)
}