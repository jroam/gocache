package file

import (
	"fmt"
	"github.com/jroam/gocache/util"
	"os"
	"path/filepath"
	"time"
)

type Cache struct {
	Path string
}

func New() *Cache {
	return &Cache{
		Path: "./cache_dir",
	}
}

func (c *Cache) Get(key interface{}) interface{} {
	rsp, err := os.ReadFile(c.Path + "/" + util.Sha(key.(string)))
	if err != nil {
		return nil
	}
	if string(rsp) == "<nil>" {
		return nil
	}
	return string(rsp)
}

func (c *Cache) Set(key, value interface{}, expire ...int64) error {
	k, ok := key.(string)
	if !ok {
		return fmt.Errorf("invalid key type, expected string")
	}

	// 使用filepath.Join确保路径的安全拼接
	filePath := filepath.Join(c.Path, util.Sha(k))

	// 创建目录，确保存在
	err := os.MkdirAll(c.Path, 0755) // 使用更合理的权限设置
	if err != nil {
		return err
	}
	var now int64
	if expire != nil && len(expire) > 0 {
		now = time.Now().Unix() + expire[0]
	} else {
		now = -1
	}

	// 将值转换为字符串并写入文件
	err = os.WriteFile(filePath, []byte(fmt.Sprintf("%d:%v", now, value)), 0644) // 使用更合理的文件权限设置
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) SetWithExpire(key, value interface{}, expiration time.Duration) error {
	return nil
}

func (c *Cache) Delete(key interface{}) error {
	err := os.Remove(c.Path + "/" + util.Sha(key.(string)))
	if err != nil {
		return err
	}
	return nil
}

//
// func (c *Cache) getType(key interface{}) string {
// 	switch key.(type) {
// 	case string:
// 		return "string"
// 	case int:
// 		return "int"
// 	case int64:
// 		return "int64"
// 	case int32:
// 		return "int32"
// 	case int16:
// 		return "int16"
// 	case int8:
// 		return "int8"
// 	case uint:
// 		return "uint"
// 	case uint64:
// 		return "uint64"
// 	case uint32:
// 		return "uint32"
// 	case uint16:
// 		return "uint16"
// 	case uint8:
// 		return "uint8"
// 	case float64:
// 		return "float64"
// 	case float32:
// 		return "float32"
// 	case bool:
// 		return "bool"
// 	case time.Time:
// 		return "time.Time"
// 	case time.Duration:
// 		return "time.Duration"
// 	case []byte:
// 		return "[]byte"
// 	case []string:
// 		return "[]string"
// 	case []int:
// 		return "[]int"
// 	case []int64:
// 		return "[]int64"
// 	case []int32:
// 		return "[]int32"
// 	case []int16:
// 		return "[]int16"
// 	case []int8:
// 		return "[]int8"
// 	case []uint:
// 		return "[]uint"
// 	case []uint64:
// 		return "[]uint64"
// 	case []uint32:
// 		return "[]uint32"
// 	case []uint16:
// 		return "[]uint16"
// 	case []float64:
// 		return "[]float64"
// 	case []float32:
// 		return "[]float32"
//
// 	default:
// 		return "interface{}"
// 	}
// }
