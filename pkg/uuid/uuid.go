package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

type UuidImpl struct {
	value string
}

func (u UuidImpl) NewV4() (string, error) {
	return newV4()
}

func NewUuid(uuid string) *UuidImpl {
	return &UuidImpl{value: uuid}
}

type Uuid [16]byte

func (uuid Uuid) String() string {
	buf := make([]byte, 36)
	hex.Encode(buf[:8], uuid[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])
	return string(buf)
}

func newV4() (string, error) {
	var uuid Uuid

	if _, err := io.ReadFull(rand.Reader, uuid[:]); err != nil {
		return uuid.String(), err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80
	return uuid.String(), nil
}
