package xtea

import (
	"encoding/binary"
	"errors"
)

var (
	key    = []byte{'c', 'a', 'v', 'a', 'n', '.', 'x', 'u', '@', 'g', 'i', 't', 'h', 'u', 'b', '.', 'c', 'o', 'm'}
	cipher [4]uint32
)

func DecodeUint32(buff []byte, pos int) (uint32, int, error) {
	if pos+4 > len(buff) {
		return 0, pos, errors.New("decode buff is not enough")
	}

	res := binary.BigEndian.Uint32(buff[pos:])
	return res, pos + 4, nil
}

func init() {
	pos := 0
	cipher[0], pos, _ = DecodeUint32(key, pos)
	cipher[1], pos, _ = DecodeUint32(key, pos)
	cipher[2], pos, _ = DecodeUint32(key, pos)
	cipher[3], pos, _ = DecodeUint32(key, pos)
}
