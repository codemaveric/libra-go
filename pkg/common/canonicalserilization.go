package common

import (
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
	//Need to work on this
	c.Position += num
	return c.Data[c.Position:num]
}
