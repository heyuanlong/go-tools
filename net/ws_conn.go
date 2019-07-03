package net

import (
	"sync"
	"time"
	//"net"
	"github.com/gorilla/websocket"

	klog "log"
)

const (
	TextMessage   = websocket.TextMessage
	BinaryMessage = websocket.BinaryMessage
)

type WSConnStruct struct {
	procotolType int
	sconn        *websocket.Conn
	agent        AgentInterface

	sendChan    chan []byte
	isCloseSend bool
	mutex       sync.Mutex

	readTimeOut int
	messageType int
}

func NewWSConnStruct(sconn *websocket.Conn, a AgentInterface, readTimeOut int, messageType int) *WSConnStruct {
	return &WSConnStruct{
		procotolType: PROCOTOL_WEBSOCKET_TYPE,
		sconn:        sconn,
		agent:        a,
		sendChan:     make(chan []byte, 1024),
		readTimeOut:  readTimeOut,
		messageType:  messageType,
	}
}

func (this *WSConnStruct) Close() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.isCloseSend == false {
		close(this.sendChan)
		this.isCloseSend = true
	}
}
func (this *WSConnStruct) IsClose() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.isCloseSend
}

func (this *WSConnStruct) Send(msg []byte) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.isCloseSend == false {
		this.sendChan <- msg
	}
}

func (this *WSConnStruct) Run() {
	go this.read()
	this.write()
}
func (this *WSConnStruct) read() {
	defer func() {
		this.Close()
		this.agent.OnClose(this)
		klog.Println("go read out")
	}()
	for {
		this.sconn.SetReadDeadline(time.Now().Add(time.Duration(this.readTimeOut) * time.Millisecond))
		_, msg, err := this.sconn.ReadMessage()
		//klog.Println("mt:",mt)
		if err != nil {
			klog.Println("ReadMessage fail:", err)
			return
		}

		this.agent.OnMessage(this, msg[0:])
	}
}
func (this *WSConnStruct) write() {
	defer func() {
		this.sconn.Close()
		klog.Println("go write out")
	}()
	for {
		select {
		case message, ok := <-this.sendChan:
			if !ok {
				this.sconn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := this.sconn.WriteMessage(this.messageType, message)
			if err != nil {
				klog.Println("WriteMessage fail ", err)
				return
			}
		}
	}
}
