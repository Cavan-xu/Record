package common

import (
	"encoding/binary"
	"errors"
)

func EncodeUint16(i uint16) []byte {
	buff := make([]byte, 2)
	binary.BigEndian.PutUint16(buff, i)
	return buff
}

func DecodeUint16(buff []byte, pos int) (uint16, int, error) {
	if pos+2 > len(buff) {
		return 0, pos, errors.New("decode buff is not enough")
	}

	res := binary.BigEndian.Uint16(buff[pos:])
	return res, pos + 2, nil
}

func EncodeUint32(i uint32) []byte {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, i)
	return buff
}

func DecodeUint32(buff []byte, pos int) (uint32, int, error) {
	if pos+4 > len(buff) {
		return 0, pos, errors.New("decode buff is not enough")
	}

	res := binary.BigEndian.Uint32(buff[pos:])
	return res, pos + 4, nil
}

// 大端序, 低地址放字节高位
func EncodeInt32(i int32) []byte {
	buff := make([]byte, 0, 4)
	return append(buff, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func DecodeInt32(buff []byte, pos int) (int32, int, error) {
	if pos+4 > len(buff) {
		return 0, 0, errors.New("decode buff is not enough")
	}

	res := binary.BigEndian.Uint32(buff[pos:])
	return int32(res), pos + 4, nil
}

func EncodeInt64(i int64) []byte {
	buff := make([]byte, 0, 8)
	return append(buff, byte(i>>56), byte(i>>48), byte(i>>40), byte(i>>32), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func DecodeInt64(buff []byte, pos int) (int64, int, error) {
	if pos+8 > len(buff) {
		return 0, pos, errors.New("decode buff is not enough")
	}

	res := binary.BigEndian.Uint64(buff[pos:])
	return int64(res), pos + 8, nil
}

func EncodeString(s string) []byte {
	buff := EncodeUint16(uint16(len(s)))
	return append(buff, []byte(s)...)
}

func DecodeString(buff []byte, pos int) (string, int, error) {
	strLength, pos, err := DecodeUint16(buff, pos)
	if err != nil {
		return "", pos, err
	}

	length := int(strLength)
	if length+pos > len(buff) {
		return "", 0, errors.New("decode buff is not enough")
	}

	tmp := make([]byte, length)
	copy(tmp, buff[pos:])
	return string(tmp), pos + length, nil
}
