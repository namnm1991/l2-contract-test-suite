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

func Uint48ToBytes(i uint64) []byte{
	var bur bytes.Buffer
	err := binary.Write(&bur, binary.BigEndian, &i)
	if err != nil {
		panic("Uint48ToBytes")
	}
	return bur.Bytes()[2:]
}

func Uint8ToByte(i uint8) byte {
	return byte(i)
}

func Uint16ToByte(i uint16) []byte {
	var bur bytes.Buffer
	err := binary.Write(&bur, binary.BigEndian, &i)
	if err != nil {
		panic("Uint16ToByte")
	}
	return bur.Bytes()
}
