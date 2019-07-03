package main

import (

	knet "github.com/heyuanlong/go-utils/net"
	klog "log"
)

type AgentSruct struct {
}

func NewAgentSruct( ) *AgentSruct  {
	return &AgentSruct{
	}
}

func (this *AgentSruct) OnConnecting() error {
	klog.Println("OnConnecting")
	return nil
}
func (this *AgentSruct) OnConnected(clientConn knet.ConnInterface) error {
	klog.Println("OnConnected")
	return nil
}
func (this *AgentSruct) OnMessage(clientConn knet.ConnInterface, msg []byte) error {
	klog.Println("OnMessage:",clientConn,msg)
	clientConn.Send(msg)
	c ,ok := clientConn.(*knet.TCPConnStruct)
	if ok {
		c.SendSync(msg)
		c.Write(msg)
	}
	return nil
}
func (this *AgentSruct) OnClose(clientConn knet.ConnInterface) error {
	klog.Println("OnClose")
	return nil
}
func (this *AgentSruct) OnError(clientConn knet.ConnInterface) error {
	klog.Println("OnError")
	return nil
}
func (this *AgentSruct) OnTimeOut(clientConn knet.ConnInterface) error {
	klog.Println("OnTimeOut")
	return nil
}
func (this *AgentSruct) CheckPackage(msg []byte) int {
	klog.Println("CheckPackage")
	return len(msg)
}


func main() {
	a := NewAgentSruct()
	s :=knet.NewTCPServer(":8081",a,2000)
	s.Run()
}