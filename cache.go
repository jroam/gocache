package gocache

import (
	"github.com/jroam/gocache/driver/file"
	"time"
)

type Cache interface {
	Set(key, value interface{}) error
	SetWithExpire(key, value interface{}, expiration time.Duration) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
}

var drivers = map[string]Cache{
	"file": file.New(),
}

//	func New(driver string) Cache {
//		return drivers[driver]
//	}
func Set(key, value interface{}) error {
	return drivers["file"].Set(key, value)
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
