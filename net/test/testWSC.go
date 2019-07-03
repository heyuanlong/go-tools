package main

import (
	klog "log"
	"time"

	knet "github.com/heyuanlong/go-utils/net"
)

type AgentSruct struct {
}

func NewAgentSruct() *AgentSruct {
	return &AgentSruct{}
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
	klog.Println("OnMessage:", clientConn, string(msg))
	clientConn.Send(msg)
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
	s := knet.NewWSClient("ws", "127.0.0.1:8081", "/ws", nil, a, 100000, knet.TextMessage)
	c := s.Run()
	c.Send([]byte("connect ok"))
	time.Sleep(55555555555555)
}
