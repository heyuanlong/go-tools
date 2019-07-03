package net

import (
	klog "log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	scheme      string
	addr        string
	path        string
	header      http.Header
	agent       AgentInterface
	readTimeOut int
	messageType int
}

func NewWSClient(scheme, addr, path string, header http.Header, a AgentInterface, readTimeOutMillisecond int, messageType int) *WSClient {
	return &WSClient{
		scheme:      scheme,
		addr:        addr,
		path:        path,
		header:      header,
		agent:       a,
		readTimeOut: readTimeOutMillisecond,
		messageType: messageType,
	}
}

func (this *WSClient) Run() ConnInterface {

	u := url.URL{Scheme: this.scheme, Host: this.addr, Path: this.path}
	klog.Println("connecting to ", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), this.header)
	if err != nil {
		klog.Println("dial:", err)
		return nil
	}

	var conn ConnInterface
	conn = NewWSConnStruct(c, this.agent, this.readTimeOut, this.messageType)
	err = this.agent.OnConnected(conn)
	if err != nil {
		klog.Println("OnConnected fail:", err)
		return nil
	}
	go conn.Run()
	return conn
}
