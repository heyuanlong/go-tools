package net

import (
	"net"
	"sync"
	"time"

	klog "log"
)

type TCPConnStruct struct {
	procotolType int
	sconn        net.Conn
	agent        AgentInterface

	sendChan    chan []byte
	isCloseSend bool
	mutex       sync.Mutex

	readTimeOut int

	writeBuffer    []byte
	writeBufferMax int
	writeTimeOut   int
}

func NewTCPConnStruct(sconn net.Conn, a AgentInterface, readTimeOut int) *TCPConnStruct {
	return &TCPConnStruct{
		procotolType:   PROCOTOL_TCP_TYPE,
		sconn:          sconn,
		agent:          a,
		sendChan:       make(chan []byte, 1024),
		readTimeOut:    readTimeOut,
		writeBuffer:    make([]byte, 0),
		writeBufferMax: 1024,
		writeTimeOut:   1000,
	}
}

func (this *TCPConnStruct) Close() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.isCloseSend == false {
		close(this.sendChan)
		this.isCloseSend = true
	}
}
func (this *TCPConnStruct) IsClose() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.isCloseSend
}

func (this *TCPConnStruct) Send(msg []byte) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.isCloseSend == false {
		this.sendChan <- msg
	}
}
func (this *TCPConnStruct) SendSync(msg []byte) {
	if len(this.writeBuffer) == 0 {
		this.sconn.SetWriteDeadline(time.Now().Add(time.Duration(this.writeTimeOut) * time.Millisecond))
		n, err := this.sconn.Write(msg)
		if err != nil {
			this.sconn.Close()
			klog.Println("WriteMessage fail :", err)
			return
		}
		if len(msg) > n {
			this.writeBuffer = append(this.writeBuffer, msg[n:]...)
		}
	} else {
		if len(this.writeBuffer) > this.writeBufferMax {
			this.sconn.Close()
			klog.Println("writeBufferMax")
			return
		}
		this.writeBuffer = append(this.writeBuffer, msg[:]...)
		this.sconn.SetWriteDeadline(time.Now().Add(time.Duration(this.writeTimeOut) * time.Millisecond))
		n, err := this.sconn.Write(this.writeBuffer)
		if err != nil {
			this.sconn.Close()
			klog.Println("WriteMessage fail :", err)
			return
		}
		if len(this.writeBuffer) > n {
			this.writeBuffer = this.writeBuffer[n:]
		}
	}
	return
}
func (this *TCPConnStruct) Write(msg []byte) (n int, err error) {
	this.sconn.SetWriteDeadline(time.Now().Add(time.Duration(this.writeTimeOut) * time.Millisecond))
	return this.sconn.Write(msg)
}

func (this *TCPConnStruct) Run() {
	go this.read()
	this.write()
}

func (this *TCPConnStruct) read() {
	defer func() {
		this.Close()
		this.agent.OnClose(this)
		klog.Println("go read out")
	}()
	var bufBuf = make([]byte, 0)
	var msgBuf = make([]byte, READ_MSG_SIZE_MAX/10)
	for {
		this.sconn.SetReadDeadline(time.Now().Add(time.Duration(this.readTimeOut) * time.Millisecond))
		n, err := this.sconn.Read(msgBuf)
		if err != nil {
			if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
				this.agent.OnTimeOut(this)
				continue
			} else {
				klog.Println("read close or fail")
				return
			}
		}
		if (len(bufBuf) + n) > READ_MSG_SIZE_MAX {
			klog.Println("buf too big")
			return
		}
		//klog.Println(msgBuf)
		bufBuf = append(bufBuf, msgBuf[0:n]...)
		for {
			packageLen := this.agent.CheckPackage(bufBuf)
			if packageLen != 0 {
				this.agent.OnMessage(this, bufBuf[0:packageLen])
				bufBuf = bufBuf[packageLen:]
			} else {
				break
			}
		}

	}
}
func (this *TCPConnStruct) write() {
	defer func() {
		this.sconn.Close()
		klog.Println("go write out")
	}()
	for {
		select {
		case message, ok := <-this.sendChan:
			if !ok {
				return
			}
			_, err := this.sconn.Write(message)
			if err != nil {
				klog.Println("WriteMessage fail :", err)
				return
			}
		}
	}
}
