package xtea

/*
	xTea: tea 算法优化版本，一种高效的加密解密算法
*/

import (
	"math/rand"
	"sync"
	"unsafe"

	"awesomeProject/practical_skills/common"
)

var (
	once sync.Once
	xTea *XTea

	key = []byte{'c', 'a', 'v', 'a', 'n', '.', 'x', 'u', '@', 'g', 'i', 't', 'h', 'u', 'b', '.', 'c', 'o', 'm'}
)

type XTea struct {
	cipher    [4]uint32
	calcCount int32
}

func NewXTea(calcCount int32) *XTea {
	once.Do(func() {
		var (
			pos    = 0
			cipher [4]uint32
		)

		cipher[0], pos, _ = common.DecodeUint32(key, pos)
		cipher[1], pos, _ = common.DecodeUint32(key, pos)
		cipher[2], pos, _ = common.DecodeUint32(key, pos)
		cipher[3], pos, _ = common.DecodeUint32(key, pos)

		xTea = &XTea{
			cipher:    cipher,
			calcCount: calcCount,
		}
	})

	return xTea
}

func (t XTea) xTea(v [8]byte, n int32) (o [8]byte) {
	y := *(*uint32)(unsafe.Pointer(&v))
	z := *(*uint32)(unsafe.Pointer(&v[4]))
	delta := uint32(0x9e3779b9)
	if n > 0 {
		limit := delta * uint32(n)
		sum := uint32(0)
		for limit != sum {
			y += (z<<4 ^ z>>5) + z ^ sum + t.cipher[sum&3]
			sum += delta
			z += (y<<4 ^ y>>5) + y ^ sum + t.cipher[sum>>11&3]
		}
	} else {
		sum := delta * uint32(-n)
		for sum != 0 {
			z -= (y<<4 ^ y>>5) + y ^ sum + t.cipher[sum>>11&3]
			sum -= delta
			y -= (z<<4 ^ z>>5) + z ^ sum + t.cipher[sum&3]
		}
	}

	*(*uint32)(unsafe.Pointer(&o)) = y
	*(*uint32)(unsafe.Pointer(&o[4])) = z
	return
}

func (t *XTea) Encrypt(in []byte) []byte {
	var (
		x [8]byte
		y [8]byte
		z [8]byte
	)

	*(*uint32)(unsafe.Pointer(&x)) = rand.Uint32()
	*(*uint32)(unsafe.Pointer(&x[4])) = uint32(len(in))

	buff := make([]byte, 0)
	res := make([]byte, 0)
	buff = append(buff, x[:]...)
	buff = append(buff, in...)
	buff = append(buff, y[:7]...)

	for i := 0; i+8 <= len(buff); i += 8 {
		copy(z[:], buff[i:i+8])
		o := t.xTea(z, t.calcCount)
		if i >= 8 {
			for j := 0; j < 8; j++ {
				o[j] = o[j] ^ res[i-8+j]
			}
		}
		res = append(res, o[:]...)
	}

	return res
}

func (t *XTea) Decrypt(in []byte) []byte {
	var z [8]byte
	res := make([]byte, 0)
	for i := 0; i+8 <= len(in); i += 8 {
		copy(z[:], in[i:i+8])
		if i >= 8 {
			for j := 0; j < 8; j++ {
				z[j] = z[j] ^ in[i-8+j]
			}
		}
		o := t.xTea(z, -t.calcCount)
		res = append(res, o[:]...)
	}

	if len(res) > 8 {
		l := *(*uint32)(unsafe.Pointer(&res[4]))
		if 8+int(l) <= len(res) {
			return res[8 : 8+l]
		}
	}

	return res[0:0]
}
