package main

import (
	"fmt"
	"time"

	"github.com/heyuanlong/go-tools/runtime"
)

func main() {
	runtime.CreateDir("log")
	runtime.Pid("")

	runtime.MainGetPanic(func() {
		for index := 0; index < 10; index++ {
			time.Sleep(1 * time.Second)
			fmt.Println("MainGetPanic", index)
			if index == 9 {
				panic("")
			}
		}
	})
	runtime.MainGetPanicAndLoop(func() {
		for index := 0; index < 10; index++ {
			time.Sleep(1 * time.Second)
			fmt.Println("MainGetPanicAndLoop", index)
			if index == 9 {
				panic("")
			}
		}
	})
}
