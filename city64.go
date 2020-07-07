package cityhash

import (
	"encoding/binary"
	"hash"
)

// City64 ...
type City64 struct {
	s []byte
}

var _ hash.Hash64 = (*City64)(nil)
var _ hash.Hash = (*City64)(nil)

// New64 ...
func New64() hash.Hash64 {
	return &City64{}
}

// Sum ...
func (c *City64) Sum(b []byte) []byte {
	b2 := make([]byte, 8)
	binary.BigEndian.PutUint64(b2, c.Sum64())
	b = append(b, b2...)
	return b
}

// Sum64 ...
func (c *City64) Sum64() uint64 {
	return CityHash64(c.s, uint32(len(c.s)))
}

// Reset ...
func (c *City64) Reset() {
	c.s = c.s[0:0]
}

// BlockSize ...
func (c *City64) BlockSize() int {
	return 1
}

func (c *City64) Write(s []byte) (n int, err error) {
	c.s = append(c.s, s...)
	return len(s), nil
}

// Size 8
func (c *City64) Size() int {
	return 8
}
