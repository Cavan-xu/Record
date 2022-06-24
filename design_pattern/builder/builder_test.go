package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConnectionPool(t *testing.T) {
	pool1, err := NewConnectionPool("pool1")
	assert.Nil(t, err)
	assert.Equal(t, pool1.maxConnect, 10)
	assert.Equal(t, pool1.maxHandler, 10)

	pool2, err := NewConnectionPool("pool2", WithMaxConnect(1))
	assert.Nil(t, err)
	assert.Equal(t, pool2.maxConnect, 1)
	assert.Equal(t, pool2.maxHandler, 10)

	pool3, err := NewConnectionPool("pool2", WithMaxConnect(1), WithMaxHandler(2))
	assert.Nil(t, err)
	assert.Equal(t, pool3.maxConnect, 1)
	assert.Equal(t, pool3.maxHandler, 2)
}
