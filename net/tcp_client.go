package net

import (
	"net"
	"time"

	klog "log"
)


type TCPClient struct {
	addr 				string
	agent 				AgentInterface
	readTimeOut 		int
}

func NewTCPClient(addr string, a AgentInterface,readTimeOutMillisecond int)*TCPClient {
	clent := &TCPClient{
		addr:addr,
		agent:a,
		readTimeOut:readTimeOutMillisecond,
	}
	return clent
}

func (ts *TCPClient) Run() ConnInterface {

	tcpconn, err := net.DialTimeout("tcp",ts.addr, time.Duration(ts.readTimeOut) * time.Millisecond)
	if err != nil {
		if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
			klog.Println("connect timeout:",ts.addr)
		} else {
			klog.Println("connect fail:",err,ts.addr)
		}
		return nil
	}

	var conn ConnInterface
	conn = NewTCPConnStruct(tcpconn,ts.agent,ts.readTimeOut)
	err = ts.agent.OnConnected(conn)
	if err != nil{
		klog.Println("OnConnected fail:",err)
		return nil
	}
	go conn.Run()
	return conn
}
