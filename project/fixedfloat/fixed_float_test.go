package fixedfloat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToFixedFloat(t *testing.T) {
	assert.Equal(t, 1<<enLargeBit, IntToFixedFloat(1).Int())
	assert.Equal(t, 100<<enLargeBit, IntToFixedFloat(100).Int())
}

func TestStrToFixedFloat(t *testing.T) {
	assert.Equal(t, 100<<enLargeBit, StrToFixedFloat("100").Int())
	assert.Equal(t, float32(1000.0001), StrToFixedFloat("1000.0001").Float32())
	assert.Equal(t, 1000.0000915527344, StrToFixedFloat("1000.0001").Float64())
}
