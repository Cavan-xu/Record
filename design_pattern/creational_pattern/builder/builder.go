package builder

import "errors"

/*
	建造者模式: 一般用于创建属性较多的对象，且该结构体的属性有约束条件，存在必选和菲比选的条件，一般用于创建不可变的对象
*/

type ConnectionPool struct {
	name       string
	maxConnect int
	maxHandler int
}

type ConnectionPoolOption struct {
	maxConnect int
	maxHandler int
}

type ConnectionPoolOptionFunc func(option *ConnectionPoolOption)

func WithMaxConnect(maxConnect int) func(option *ConnectionPoolOption) {
	return func(option *ConnectionPoolOption) {
		option.maxConnect = maxConnect
	}
}

func WithMaxHandler(maxHandler int) func(option *ConnectionPoolOption) {
	return func(option *ConnectionPoolOption) {
		option.maxHandler = maxHandler
	}
}

func NewConnectionPool(name string, options ...ConnectionPoolOptionFunc) (*ConnectionPool, error) {
	if name == "" {
		return nil, errors.New("name can not be null")
	}

	option := &ConnectionPoolOption{
		maxHandler: 10,
		maxConnect: 10,
	}
	for _, fn := range options {
		fn(option)
	}

	if option.maxConnect <= 0 {
		return nil, errors.New("maxConnect must be greater than 0")
	}

	if option.maxHandler <= 0 {
		return nil, errors.New("maxHandler must be greater than 0")
	}

	return &ConnectionPool{
		name:       name,
		maxConnect: option.maxConnect,
		maxHandler: option.maxHandler,
	}, nil
}
