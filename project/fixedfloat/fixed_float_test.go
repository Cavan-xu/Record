package fixedfloat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToFixedFloat(t *testing.T) {
	assert.Equal(t, 1, IntToFixedFloat(1).Int())
	assert.Equal(t, 100, IntToFixedFloat(100).Int())
}

func TestStrToFixedFloat(t *testing.T) {
	assert.Equal(t, 100, StrToFixedFloat("100").Int())
	assert.Equal(t, float32(1000.0001), StrToFixedFloat("1000.0001").Float32())
	assert.Equal(t, 1000.0000915527344, StrToFixedFloat("1000.0001").Float64())
}

func TestFixedFloat_Ceil(t *testing.T) {
	assert.Equal(t, 1, IntToFixedFloat(1).Ceil())
	assert.Equal(t, 1, StrToFixedFloat("0.1").Ceil())
}

func TestFixedFloat_Floor(t *testing.T) {
	assert.Equal(t, 1, IntToFixedFloat(1).Floor())
	assert.Equal(t, 1, StrToFixedFloat("1.1").Floor())
}

func TestFixedFloat_Round(t *testing.T) {
	assert.Equal(t, 1, IntToFixedFloat(1).Round())
	assert.Equal(t, 1, StrToFixedFloat("1.4").Round())
	assert.Equal(t, 2, StrToFixedFloat("1.5").Round())
}

func TestFixedFloat_Add(t *testing.T) {
	assert.Equal(t, 2, IntToFixedFloat(1).Add(IntToFixedFloat(1)).Int())
	assert.Equal(t, 2, StrToFixedFloat("1.1").Add(StrToFixedFloat("1.2")).Int())

	t.Logf("4.9 + 0.1 = %f", StrToFixedFloat("4.9").Add(StrToFixedFloat("0.1")).Float64())
	t.Logf("4.9 + (-0.1) = %f", StrToFixedFloat("4.9").Add(StrToFixedFloat("-0.1")).Float64())
}

func TestFixedFloat_Sub(t *testing.T) {
	assert.Equal(t, 1, IntToFixedFloat(2).Sub(IntToFixedFloat(1)).Int())

	t.Logf("4.4 - 2.1 = %f", StrToFixedFloat("4.4").Sub(StrToFixedFloat("2.1")).Float64())
	t.Logf("4.4 - (-2.1) = %f", StrToFixedFloat("4.4").Sub(StrToFixedFloat("-2.1")).Float64())
}

func TestFixedFloat_Mul(t *testing.T) {
	assert.Equal(t, 2, IntToFixedFloat(1).Mul(IntToFixedFloat(2)).Int())

	t.Logf("4 * 1.1 = %f", StrToFixedFloat("4").Mul(StrToFixedFloat("1.1")).Float64())
	t.Logf("4 * (-1.1) = %f", StrToFixedFloat("4").Mul(StrToFixedFloat("-1.1")).Float64())
}

func TestFixedFloat_Div(t *testing.T) {
	assert.Equal(t, 2, IntToFixedFloat(4).Div(IntToFixedFloat(2)).Int())

	t.Logf("4.2 / 2.1 = %f", StrToFixedFloat("4.2").Div(StrToFixedFloat("2.1")).Float64())
	t.Logf("4.2 / (-2.1) = %f", StrToFixedFloat("4.2").Div(StrToFixedFloat("-2.1")).Float64())
}
