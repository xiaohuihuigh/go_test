package main

import "fmt"

type IpcClient struct {
	conn chan string
}

type Request struct {
	Methond string `json:"method,string"`
	Params  Param  `json:"params,Param"`
}
type Param struct {
	K string `json:"k,string"`
	V string `json:"v,string"`
}

func (client *IPcClient) Call(request *Request) *Response {
	ReqStr := Req2Str(request)
}
