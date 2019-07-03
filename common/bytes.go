package common

import (
	"bytes"
	"encoding/binary"
)

//字节转换成整形-小端序
func BytesToInt32Little(b []byte) (int32, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

//整形转换成字节-小端序
func Int32ToBytesLittle(n int32) ([]byte, error) {
	x := n

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return []byte{}, err
	}
	return bytesBuffer.Bytes(), nil
}

//字节转换成整形-小端序
func BytesToInt64Little(b []byte) (int64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

//整形转换成字节-小端序
func Int64ToBytesLittle(n int64) ([]byte, error) {
	x := n

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return []byte{}, err
	}
	return bytesBuffer.Bytes(), nil
}

//字节转换成整形-小端序
func BytesToFloat64Little(b []byte) (float64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x float64
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

//整形转换成字节-小端序
func Float64ToBytesLittle(n float64) ([]byte, error) {
	x := n

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return []byte{}, err
	}
	return bytesBuffer.Bytes(), nil
}
