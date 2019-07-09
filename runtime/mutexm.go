package runtime

import (
	"sync"
)

type MutexM struct {
	glockInt    sync.Mutex
	glockString sync.Mutex
	lmInt       map[int64]*sync.Mutex
	lmString    map[string]*sync.Mutex
}

func NewMutexM() *MutexM {
	return &MutexM{
		lmInt:    make(map[int64]*sync.Mutex),
		lmString: make(map[string]*sync.Mutex),
	}
}

func (ts *MutexM) LockInt(i int64) {
	ts.glockInt.Lock()
	if _, ok := ts.lmInt[i]; !ok {
		ts.lmInt[i] = &sync.Mutex{}
	}
	ts.glockInt.Unlock()

	l := ts.lmInt[i]
	l.Lock()
}
func (ts *MutexM) UnLockInt(i int64) {
	l := ts.lmInt[i]
	l.Unlock()
}

func (ts *MutexM) LockString(i string) {
	ts.glockString.Lock()
	if _, ok := ts.lmString[i]; !ok {
		ts.lmString[i] = &sync.Mutex{}
	}
	ts.glockString.Unlock()

	l := ts.lmString[i]
	l.Lock()
}
func (ts *MutexM) UnLockString(i string) {
	l := ts.lmString[i]
	l.Unlock()
}
