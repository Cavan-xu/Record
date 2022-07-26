package signature

/*
	基于 xTea 的签名算法
*/

import (
	"encoding/base64"
	"time"

	"awesomeProject/common"
	"awesomeProject/practical_skills/xtea"
)

const (
	defaultExpireTime = 24 * 60 * 60 // 默认过期时间为一天
	defaultCalcCount  = 32           // xTea 计算次数
)

type Signature struct {
	UserId     int32
	UserName   string
	ExpireTime int64
}

func NewSignature(userId int32, userName string) *Signature {
	return &Signature{
		UserId:     userId,
		UserName:   userName,
		ExpireTime: time.Now().Unix() + defaultExpireTime,
	}
}

func (s *Signature) Encrypt() string {
	buff := make([]byte, 0)
	buff = append(buff, common.EncodeInt32(s.UserId)...)
	buff = append(buff, common.EncodeString(s.UserName)...)
	buff = append(buff, common.EncodeInt64(s.ExpireTime)...)

	bytes := xtea.NewXTea(defaultCalcCount).Encrypt(buff)
	return base64.StdEncoding.EncodeToString(bytes)
}

func (s *Signature) Decrypt(str string) error {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

	var pos int

	bytes = xtea.NewXTea(defaultCalcCount).Decrypt(bytes)
	s.UserId, pos, err = common.DecodeInt32(bytes, pos)
	if err != nil {
		return err
	}
	s.UserName, pos, err = common.DecodeString(bytes, pos)
	if err != nil {
		return err
	}
	s.ExpireTime, pos, err = common.DecodeInt64(bytes, pos)
	if err != nil {
		return err
	}

	return nil
}
