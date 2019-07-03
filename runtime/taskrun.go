package runtime

import (
	"sync"
	"time"
)

type agent interface {
	DoWork(v interface{})
	TimeOut()
}

type TaskRunServer struct {
	a       agent
	runNums int64            //协程数量
	vchan   chan interface{} //chan缓存长度
	isRun   bool             //该实例是否执行中
	timeOut int64            //无任务执行时长，超时提醒
	mutex   sync.Mutex
}

func NewTaskRunServer(a agent, runNums int64, chanLens int64, timeOut int64) *TaskRunServer {
	return &TaskRunServer{
		a:       a,
		runNums: runNums,
		isRun:   false,
		timeOut: timeOut,
		vchan:   make(chan interface{}, chanLens),
	}
}

func (ts *TaskRunServer) Send(v interface{}) {
	if ts.isRun == false {
		return
	}
	ts.vchan <- v
}

func (ts *TaskRunServer) Run() {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	if ts.isRun {
		return
	}
	ts.isRun = true
	for i := int64(0); i < ts.runNums; i++ {
		go ts.work()
	}
}

func (ts *TaskRunServer) work() {
	var v interface{}
	if ts.timeOut > 0 {
		for {
			select {
			case v = <-ts.vchan:
				ts.a.DoWork(v)
			case <-time.After(time.Millisecond * time.Duration(ts.timeOut)):
				ts.a.TimeOut()
			}
		}
	} else {
		for {
			select {
			case v = <-ts.vchan:
				ts.a.DoWork(v)
			}
		}
	}
}
