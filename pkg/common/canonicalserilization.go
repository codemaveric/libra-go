package common

import (
	"bytes"
	"encoding/binary"
)

type CanonicalSerializer struct {
	Data     []byte
	Position uint64
}

func NewCanonicalSerializer(data []byte) *CanonicalSerializer {
	return &CanonicalSerializer{Data: data, Position: 0}
}

func (c *CanonicalSerializer) Read32() uint32 {
	value := binary.LittleEndian.Uint32(c.Data[c.Position:])
	c.Position += 4
	return value
}

func (c *CanonicalSerializer) Read64() uint64 {
	value := binary.LittleEndian.Uint64(c.Data[c.Position:])
	c.Position += 8
	return value
}

func (c *CanonicalSerializer) ReadXBytes(num uint64) []byte {
	r := bytes.NewReader(c.Data[c.Position:num])
	var res byteType
	binary.Read(r, binary.LittleEndian, &res)
	c.Position += num
	return res
}

type byteType []byte
