package xtea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeUint32(t *testing.T) {
	t.Log(cipher)
}

func Test_encrypt(t *testing.T) {
	in := []byte{'t', 'e', 's', 't'}
	data := encrypt(in, cipher)
	fmt.Println(data)
}

func Test_decrypt(t *testing.T) {
	in := []byte{'t', 'e', 's', 't'}
	data := encrypt(in, cipher)
	res := decrypt(data, cipher)
	assert.Equal(t, in, res)
	t.Log(string(res))
}
