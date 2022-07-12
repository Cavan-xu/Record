package common

import (
	"encoding/binary"
	"errors"
)

func DecodeUint32(buff []byte, pos int) (uint32, int, error) {
	if pos+4 > len(buff) {
		return 0, pos, errors.New("decode buff is not enough")
	}

	res := binary.BigEndian.Uint32(buff[pos:])
	return res, pos + 4, nil
}
