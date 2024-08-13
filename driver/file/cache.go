package file

import "time"

type Cache struct {
}

func New() *Cache {
	return &Cache{}
}

func (c *Cache) Get(key interface{}) interface{} {
	return "hello abel!"
}

func (c *Cache) Set(key, value interface{}) error {
	return nil
}

func (c *Cache) SetWithExpire(key, value interface{}, expiration time.Duration) error {
	return nil
}

func (c *Cache) Delete(key interface{}) error {
	return nil
}
