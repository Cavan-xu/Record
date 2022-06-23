package factory

import (
	"encoding/json"
	"encoding/xml"
)

/*
	简单工厂模式：实现简单，有较多的 if 分支，适合不会频繁改动的场景
*/

type ConfigParse interface {
	parse([]byte) error
}

type JsonConfigParse struct{}

func (j JsonConfigParse) parse(body []byte) error {
	return json.Unmarshal(body, j)
}

type XmlConfigParser struct{}

func (x XmlConfigParser) parse(body []byte) error {
	return xml.Unmarshal(body, x)
}

func NewConfigParse(c string) ConfigParse {
	switch c {
	case "json":
		return JsonConfigParse{}
	case "xml":
		return XmlConfigParser{}
	default:
		return nil
	}
}
