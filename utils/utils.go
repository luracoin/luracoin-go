package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func HeigthToString(heigth uint32) []byte {
	// 00000000
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, heigth)
	return bs
}

func BytesToInt(b []byte) uint32 {
	data := binary.LittleEndian.Uint32(b)
	return data
}
