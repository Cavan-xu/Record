package signature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignature_Encrypt(t *testing.T) {
	signature := NewSignature(999999, "cavan.xu")
	str := signature.Encrypt()
	t.Log(str)
	assert.NotEqual(t, "", str)
}

func TestSignature_Decrypt(t *testing.T) {
	signature := NewSignature(999999, "cavan.xu")
	str := signature.Encrypt()

	signature2 := &Signature{}
	err := signature2.Decrypt(str)
	assert.Nil(t, err)
	assert.Equal(t, signature, signature2)
}
