package singleton

import "sync"

/*
	单例模式：最简单的设计模式
	一个对象只创建一个实例，减少内存开销
	第一种方式：通过在 init 函数中初始化对象，能够保证只执行一次
	第二种方式：通过 sync 包下面提供的 once 能力实现单例模式
*/

type client struct{}

var (
	once      sync.Once
	_instance *client
)

func init() {
	_instance = &client{}
}

func New() *client {
	return _instance
}

func Singleton() *client {
	once.Do(func() {
		_instance = &client{}
	})

	return _instance
}
