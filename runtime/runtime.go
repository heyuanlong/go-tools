package runtime

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
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
	f, err := os.OpenFile(pidFile, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("OpenFile error:", err)
		return
	}
	defer func() {
		_ = f.Close()
	}()
	_, _ = f.WriteString(fmt.Sprintf("%d", os.Getpid()))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}
