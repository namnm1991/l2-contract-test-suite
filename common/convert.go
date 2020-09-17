package common

import (
	"bytes"
	"encoding/binary"
)

func Uint32ToBytes(i uint32) []byte {
	var bur bytes.Buffer
	err := binary.Write(&bur, binary.BigEndian, &i)
	if err != nil {
		panic("Uint32ToBytes")
	}
	return bur.Bytes()
}

func Uint8ToByte(i uint8) byte {
	return byte(i)
}
