package main

import (
	"fmt"
	"strconv"

	"github.com/heyuanlong/go-tools/log"
)

func main() {
	klog, err := log.NewLlogFile("log.log", "Debug", log.LstdFlags|log.Lshortfile, log.LOG_LEVEL_DEBUG, log.LOG_LEVEL_DEBUG, 15)
	if err != nil {
		fmt.Println(err)
		return
	}
	for index := 0; index < 1000000; index++ {
		klog.Println("12345678iowetuiosdfghjk---", strconv.Itoa(index))
	}
}
