package factory

/*
	工厂方法：当对象的创建逻辑比较复杂，，不是简单的 new 就行了，需要组和其他对象，做各种初始化操作，推荐使用工厂方法，将复杂的创造逻辑
			 拆分到各个工厂类中，让每个工厂类不至于太复杂
*/

type ConfigParseFactory interface {
	NewConfigParse() ConfigParse
}

type JsonConfigParseFactory struct{}

func (j JsonConfigParseFactory) NewConfigParse() ConfigParse {
	return JsonConfigParse{}
}

type XmlConfigParseFactory struct{}

func (x XmlConfigParseFactory) NewConfigParse() ConfigParse {
	return XmlConfigParser{}
}

func NewConfigParseFactory(c string) ConfigParseFactory {
	switch c {
	case "json":
		return JsonConfigParseFactory{}
	case "xml":
		return XmlConfigParseFactory{}
	default:
		return nil
	}
}
