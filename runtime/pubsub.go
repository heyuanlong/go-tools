package runtime

import "sync"

type subCallBackI interface {
	subCallBack(string, interface{})
}

type dataStruct struct {
	topic string
	v     interface{}
}

type pubsub struct {
	subObj []subCallBackI
	vchan  chan dataStruct //chan缓存长度
	mutex  sync.Mutex
}

func NewPubSub(subObj []subCallBackI, chanLens int64) *pubsub {
	ts := &pubsub{
		subObj: subObj,

		vchan: make(chan dataStruct, chanLens),
	}
	ts.UniqueSubObj()

	go ts.run()
	return ts
}
func (ts *pubsub) AddSub(subObj []subCallBackI) {
	ts.subObj = append(ts.subObj, subObj...)
	ts.UniqueSubObj()
}

func (ts *pubsub) Pub(topic string, v interface{}) {
	ts.vchan <- dataStruct{
		topic,
		v,
	}
}

func (ts *pubsub) run() {
	for {
		select {
		case v := <-ts.vchan:
			ts.mutex.Lock()
			defer ts.mutex.Unlock()

			for i := 0; i < len(ts.subObj); i++ {
				ts.subObj[i].subCallBack(v.topic, v.v)
			}
		default:

		}
	}
}

//去重
func (ts *pubsub) UniqueSubObj() {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	for i := 0; i < len(ts.subObj)-1; i++ {
		for j := i + 1; j < len(ts.subObj); {
			if ts.subObj[i] == ts.subObj[j] {
				ts.subObj = append(ts.subObj[:j], ts.subObj[j+1:]...)
			} else {
				j++
			}
		}
	}
}

//--------------------------------------------------------------------------------------

type pubUpAndDown struct {
	upObj       []subCallBackI
	downObj     []subCallBackI
	toUpvchan   chan dataStruct //chan缓存
	toDownvchan chan dataStruct //chan缓存
	mutex       sync.Mutex
}

func NewpubUpAndDown(upObj []subCallBackI, downObj []subCallBackI, chanLens int64) *pubUpAndDown {
	ts := &pubUpAndDown{
		upObj:   upObj,
		downObj: downObj,

		toUpvchan:   make(chan dataStruct, chanLens),
		toDownvchan: make(chan dataStruct, chanLens),
	}
	ts.UniqueObj()

	go ts.run()
	return ts
}
func (ts *pubUpAndDown) AddUpObj(upObj []subCallBackI) {
	ts.upObj = append(ts.upObj, upObj...)
	ts.UniqueObj()
}

func (ts *pubUpAndDown) AddDownObj(downObj []subCallBackI) {
	ts.downObj = append(ts.downObj, downObj...)
	ts.UniqueObj()
}

func (ts *pubUpAndDown) PubToUp(topic string, v interface{}) {
	ts.toUpvchan <- dataStruct{
		topic,
		v,
	}
}

func (ts *pubUpAndDown) PubToDown(topic string, v interface{}) {
	ts.toDownvchan <- dataStruct{
		topic,
		v,
	}
}

func (ts *pubUpAndDown) run() {
	for {
		select {
		case v := <-ts.toUpvchan:
			ts.mutex.Lock()
			defer ts.mutex.Unlock()

			for i := 0; i < len(ts.upObj); i++ {
				ts.upObj[i].subCallBack(v.topic, v.v)
			}
		case v := <-ts.toDownvchan:
			ts.mutex.Lock()
			defer ts.mutex.Unlock()

			for i := 0; i < len(ts.downObj); i++ {
				ts.downObj[i].subCallBack(v.topic, v.v)
			}
		default:

		}
	}
}

//去重
func (ts *pubUpAndDown) UniqueObj() {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	for i := 0; i < len(ts.upObj)-1; i++ {
		for j := i + 1; j < len(ts.upObj); {
			if ts.upObj[i] == ts.upObj[j] {
				ts.upObj = append(ts.upObj[:j], ts.upObj[j+1:]...)
			} else {
				j++
			}
		}
	}

	for i := 0; i < len(ts.downObj)-1; i++ {
		for j := i + 1; j < len(ts.downObj); {
			if ts.downObj[i] == ts.downObj[j] {
				ts.downObj = append(ts.downObj[:j], ts.downObj[j+1:]...)
			} else {
				j++
			}
		}
	}
}
