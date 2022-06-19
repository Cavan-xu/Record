package fixedfloat

/*
	背景：由于不同编程语言、不同机器算出来的浮点数精度可能存在误差，因此需要后端、客户端统一一个计算浮点数的方法
		 保证由后端和客户端算出来的浮点数完全一致
*/

import "strings"

const (
	enLargeBit = 16              // 十六位表示小数部分
	precision  = 1 << enLargeBit // 精度
)

type FixedFloat int64

func IntToFixedFloat(i int) FixedFloat {
	return FixedFloat(int64(i << enLargeBit))
}

func StrToFixedFloat(str string) FixedFloat {
	var (
		i               = 0
		sign            = 1
		integerPart     = 0
		decimalPart     = 0
		decimalMultiple = 1
	)

	str = strings.Trim(str, "\t\r\n")

	for ; i < len(str); i++ {
		if i == 0 && str[i] == '-' {
			sign = -1
			continue
		}
		if str[i] >= '0' && str[i] <= '9' {
			integerPart = integerPart*10 + int(str[i]-'0')
		} else {
			break
		}
	}

	value := integerPart << enLargeBit

	if i < len(str) && str[i] == '.' {
		for i = i + 1; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				decimalPart = decimalPart*10 + int(str[i]-'0')
				decimalMultiple *= 10
			}
		}
		value += decimalPart * precision / decimalMultiple // 加上的值一定小于 1 << enLargeBit，因为 decimalMultiple > decimalPart
	}

	if sign == -1 {
		value = -value
	}

	return FixedFloat(value)
}

func (f FixedFloat) Int() int {
	return int(f >> enLargeBit)
}

func (f FixedFloat) Int32() int32 {
	return int32(f >> enLargeBit)
}

func (f FixedFloat) Int64() int64 {
	return int64(f >> enLargeBit)
}

func (f FixedFloat) Float32() float32 {
	return float32(float64(f) / float64(precision))
}

func (f FixedFloat) Float64() float64 {
	return float64(f) / float64(precision)
}

func (f FixedFloat) Ceil() int {
	v := f >> enLargeBit
	if v<<enLargeBit == f {
		return int(v)
	}

	return int(v) + 1
}

func (f FixedFloat) Floor() int {
	return int(f >> enLargeBit)
}

func (f FixedFloat) Round() int {
	v := f >> enLargeBit
	if f >= v<<enLargeBit+FixedFloat(precision>>1) {
		return int(v) + 1
	}

	return int(v)
}

func (f FixedFloat) Equal(other FixedFloat) bool {
	return f == other
}

func (f FixedFloat) Less(other FixedFloat) bool {
	return f < other
}

func (f FixedFloat) LessEquals(other FixedFloat) bool {
	return f <= other
}

func (f FixedFloat) More(other FixedFloat) bool {
	return f > other
}

func (f FixedFloat) MoreEquals(other FixedFloat) bool {
	return f >= other
}

func (f FixedFloat) Neg() FixedFloat {
	return -f
}

func (f FixedFloat) Abs() FixedFloat {
	if f >= 0 {
		return f
	}

	return -f
}

func (f FixedFloat) LeftShift(shift int) FixedFloat {
	return f << shift
}

func (f FixedFloat) RightShift(shift int) FixedFloat {
	return f >> shift
}

func (f FixedFloat) Add(other FixedFloat) FixedFloat {
	return f + other
}

func (f FixedFloat) Sub(other FixedFloat) FixedFloat {
	return f - other
}

func (f FixedFloat) Mul(other FixedFloat) FixedFloat {
	integerPart := f >> enLargeBit
	decimalPart := f & FixedFloat(precision-1)

	return integerPart*other + (decimalPart*other)>>enLargeBit
}

func (f FixedFloat) Div(other FixedFloat) FixedFloat {
	return f << enLargeBit / other
}

func (f FixedFloat) Sqrt() FixedFloat {
	return 0
}
