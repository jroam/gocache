package file

import (
	"fmt"
	"github.com/jroam/gocache/util"
	"os"
	"path/filepath"
	"time"
)

type Cache struct {
	path string
}

func New() *Cache {
	return &Cache{
		path: "./cache_dir",
	}
}

func (c *Cache) Get(key interface{}) interface{} {
	rsp, err := os.ReadFile(c.path + "/" + util.Sha(key.(string)))
	if err != nil {
		return nil
	}
	return string(rsp)
}

func (c *Cache) Set(key, value interface{}) error {
	k, ok := key.(string)
	if !ok {
		return fmt.Errorf("invalid key type, expected string")
	}

	// 使用filepath.Join确保路径的安全拼接
	filePath := filepath.Join(c.path, util.Sha(k))

	// 创建目录，确保存在
	err := os.MkdirAll(c.path, 0755) // 使用更合理的权限设置
	if err != nil {
		return err
	}

	// 将值转换为字符串并写入文件
	err = os.WriteFile(filePath, []byte(fmt.Sprintf("%v", value)), 0644) // 使用更合理的文件权限设置
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) SetWithExpire(key, value interface{}, expiration time.Duration) error {
	return nil
}

func (c *Cache) Delete(key interface{}) error {
	err := os.Remove(c.path + "/" + util.Sha(key.(string)))
	if err != nil {
		return err
	}
	return nil
}
