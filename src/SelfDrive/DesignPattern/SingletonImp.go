package DesignPattern

import "sync"

var cacheInstance *cache = new(cache)
var cacheInstanceOnce *cache
var once sync.Once

type cache struct {
	Storage []interface{}
}

func init() {
	cacheInstance = new(cache) // init 方法只会执行一次，所以不需要担心重复初始化的问题
}

func GetCacheInstance() *cache {
	return cacheInstance
}

func GetCacheInstanceOnce() *cache {
	once.Do(func() {
		cacheInstanceOnce = new(cache)
	})
	return cacheInstanceOnce
}
