package net

import (
	"net"
	klog "log"
)


type TCPServer struct {
	addr 				string
	agent 				AgentInterface
	readTimeOut 		int
}
func NewTCPServer(addr string, a AgentInterface,readTimeOutMillisecond int )*TCPServer {
	return &TCPServer{
			addr:addr,
			agent:a,
			readTimeOut:readTimeOutMillisecond,
		}
}

func (this *TCPServer) Run(){

	listen_sock,err := net.Listen("tcp",this.addr)
	if err != nil{
		klog.Println("net Listen fail:",err)
	}
	defer listen_sock.Close()
	for{
		tcpconn,err := listen_sock.Accept()
		if err != nil {
			klog.Println("listen_sock.Accept error:",err)
			continue
		}
		var conn ConnInterface
		conn =NewTCPConnStruct(tcpconn,this.agent,this.readTimeOut)
		err = this.agent.OnConnected(conn)
		if err != nil{
			klog.Println("OnConnected fail:",err)
			return
		}
		go conn.Run()
	}
}