package main

import "fmt"

type IpcClient struct {
	conn chan string
}

func (client *IPcClient) Call(request *Request) *Response {
	ReqStr := Req2Str(request)
}
