package net

import (
	"net/http"

	"github.com/gorilla/websocket"

	klog "log"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type WSServer struct {
	addr 			string
	agent 			AgentInterface
	readTimeOut 	int
	messageType 	int
}
func NewWSServer(addr string, a AgentInterface,readTimeOutMillisecond 	int,messageType int)*WSServer {
	return &WSServer{
		addr:addr,
		agent:a,
		readTimeOut:readTimeOutMillisecond,
		messageType:messageType,
		}
}

func (this *WSServer) Run(){

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", this.serveWs)
	err := http.ListenAndServe(this.addr, nil)
	if err != nil {
		klog.Println("ListenAndServe fail:",err)
	}
}

func (server *WSServer) serveWs( w http.ResponseWriter, r *http.Request) {

	err  := server.agent.OnConnecting()
	if err != nil{
		klog.Println("OnConnecting fail:",err)
		return
	}

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		klog.Println("upgrader.Upgrade fail:",err)
		return
	}
	var conn ConnInterface
	conn =NewWSConnStruct(wsconn,server.agent,server.readTimeOut ,server.messageType)
	err  = server.agent.OnConnected(conn)
	if err != nil{
		klog.Println("OnConnected fail:",err)
		return
	}
	go conn.Run()
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home111.html")
}

