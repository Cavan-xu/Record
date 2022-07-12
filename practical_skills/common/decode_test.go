package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeUint32(t *testing.T) {
	var val uint32 = 5

	bytes := EncodeUint32(val)
	i, _, _ := DecodeUint32(bytes, 0)

	assert.Equal(t, val, i)
}

func TestEncodeInt32(t *testing.T) {
	var val1 int32 = 5
	var val2 int32 = -5

	bytes := EncodeInt32(val1)
	i, _, _ := DecodeInt32(bytes, 0)
	assert.Equal(t, val1, i)

	bytes = EncodeInt32(val2)
	i, _, _ = DecodeInt32(bytes, 0)
	assert.Equal(t, val2, i)
}
