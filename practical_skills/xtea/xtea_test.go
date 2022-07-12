package xtea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewXTea(t *testing.T) {
	assert.NotNil(t, NewXTea(32))
}

func Test_encrypt(t *testing.T) {
	in := []byte{'t', 'e', 's', 't'}

	xTea := NewXTea(32)
	assert.NotNil(t, xTea.Encrypt(in))
}

func Test_decrypt(t *testing.T) {
	in := []byte{'t', 'e', 's', 't'}

	xTea := NewXTea(32)
	data := xTea.Encrypt(in)
	res := xTea.Decrypt(data)

	assert.Equal(t, in, res)
	t.Log(string(res))
}
