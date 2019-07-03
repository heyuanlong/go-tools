package main

import (
	"fmt"

	"github.com/heyuanlong/go-tools/common"
)

func main() {
	fmt.Println(common.GetTimerNext(15 * 3600))
}
