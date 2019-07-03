package main

import (
	"fmt"

	"github.com/heyuanlong/go-tools/common"
)

func main() {
	gettime()
	bytes32()
	bytes64()
	xfloat64()
}

func gettime() {
	fmt.Println(common.GetTimerNext(15 * 3600))
	fmt.Println()
}

func bytes32() {
	b, _ := common.Int32ToBytesLittle(1500123789)
	n, _ := common.BytesToInt32Little(b)
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println()

	b1, _ := common.Int32ToBytesLittle(-1500123789)
	n1, _ := common.BytesToInt32Little(b1)
	fmt.Println(b1)
	fmt.Println(n1)
	fmt.Println()
}

func bytes64() {
	b, _ := common.Int64ToBytesLittle(1500123789123456789)
	n, _ := common.BytesToInt64Little(b)
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println()

	b1, _ := common.Int64ToBytesLittle(-1500123789123456789)
	n1, _ := common.BytesToInt64Little(b1)
	fmt.Println(b1)
	fmt.Println(n1)
	fmt.Println()
}

func xfloat64() {
	b, _ := common.Float64ToBytesLittle(158.9123456789)
	n, _ := common.BytesToFloat64Little(b)
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println()

	b1, _ := common.Float64ToBytesLittle(-158.9123456789)
	n1, _ := common.BytesToFloat64Little(b1)
	fmt.Println(b1)
	fmt.Println(n1)
	fmt.Println()
}
