package main

import "fmt"

type IpcClient struct {
	conn chan string
}
