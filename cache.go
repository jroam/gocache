package gocache

import (
	"github.com/jroam/gocache/driver/file"
	"time"
)

type Cache interface {
	Set(key, value interface{}, expire ...time.Duration) error
	SetWithExpire(key, value interface{}, expiration time.Duration) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
}

var entity *Cache

// 配置项
type Config struct {
	Type string `json:"type"`
}

var drivers = map[string]*Cache{
	"file": file.New(),
}

func Init(config *Config) *Cache {
	entity = drivers[config.Type]
	return entity
}

//	func New(driver string) Cache {
//		return drivers[driver]
//	}
func Set(key, value interface{}, expire ...time.Duration) error {
	return entity.Set(key, value, expire...)
}

//	func SetWithExpire(key, value interface{}, expiration time.Duration) error {
//		return file.New().SetWithExpire(key, value, expiration)
//	}
func Delete(key interface{}) error {
	return drivers["file"].Delete(key)
}

func Get(key interface{}) interface{} {
	return drivers["file"].Get(key)

}
