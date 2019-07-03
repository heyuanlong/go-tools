package net

import (

)

const (
	PROCOTOL_TCP_TYPE = 				1
	PROCOTOL_UDP_TYPE = 				2
	PROCOTOL_WEBSOCKET_TYPE = 			3

	READ_MSG_SIZE_MAX =					10240
)

type ConnInterface interface  {
	Close()
	Send(msg []byte)
	Run()
	read()
	write()
	IsClose() bool
}

type AgentInterface interface  {
	OnConnecting()(error)
	OnConnected(ConnInterface)(error)
	OnMessage(ConnInterface,[]byte)(error)
	OnClose(ConnInterface)(error)
	OnError(ConnInterface)(error)
	OnTimeOut(ConnInterface)(error)

	CheckPackage([]byte)int
}

type connManagerStruct struct {
	conns 			map[ConnInterface]int8
}