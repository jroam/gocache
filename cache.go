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

func Get(key interface{}) interface{} {
	return file.New().Get(key)

}
