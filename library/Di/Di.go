package di

import "github.com/gogf/gf/container/gmap"

var(
	globData = gmap.New(true)
)


func Set(name string, value interface{}) {
	globData.Set(name, value)
}

func Get(name string) interface{} {
	return globData.Get(name)
}

