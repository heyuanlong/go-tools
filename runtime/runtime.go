package runtime

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func MainGetPanicAndLoop(f func()) {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
				}
			}()

			f()
		}()
		time.Sleep(time.Second)
	}
}

func MainGetPanic(f func()) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	f()
}

func CreateDir(dir string) {
	if err := os.MkdirAll(dir, 755); err != nil {
		fmt.Println("MkdirAll error:", err)
		return
	}
}

func Pid(pidFile string) {
	if pidFile == "" {
		pidFile = "pid.txt"
	}
	f, err := os.OpenFile(pidFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("OpenFile error:", err)
		return
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", os.Getpid()))
}
