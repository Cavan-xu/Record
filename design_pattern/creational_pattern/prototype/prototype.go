package prototype

import (
	"encoding/json"
	"time"
)

/*
	原型模式：通过复制一个已存在的实例来返回新的实例，而不是新建实例。被复制的实例即为原型
			 例如，一个对象需要在一个高代价的数据库操作之后被创建。我们可以缓存该对象，在下一个请求时返回它的克隆，
			 在需要的时候更新数据库，以此来减少数据库调用。
*/

type Source struct {
	name       string
	val        int32
	updateTime time.Time
}

func (s *Source) Clone() *Source {
	clone := &Source{}
	bytes, _ := json.Marshal(s)
	json.Unmarshal(bytes, clone)

	return clone
}

type DataMap map[string]*Source

func (d *DataMap) Update(updateDataList []*Source) *DataMap {
	newDataMap := DataMap{}

	for _, s := range *d {
		newDataMap[s.name] = s
	}

	for _, d := range updateDataList {
		newDataMap[d.name] = d.Clone()
	}

	return &newDataMap
}
