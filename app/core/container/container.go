package container

import (
    "log"
    "strings"
    "sync"

    "component-combine/app/global/my_errors"
    "component-combine/app/global/variable"
)

// 定义一个全局键值对存储容器
var sMap sync.Map

// 定义一个容器结构体
type containers struct {
}

// CreateContainersFactory 创建一个容器工厂
func CreateContainersFactory() *containers {
    return &containers{}
}

// Set 1.以键值对的形式将代码注册到容器
func (c *containers) Set(key string, value interface{}) (res bool) {
    if _, exists := c.KeyIsExists(key); exists == false {
        sMap.Store(key, value)
        res = true
    } else {
        // 程序启动阶段,zapLog 未初始化,使用系统 log 打印启动时候发生的异常日志
        if variable.ZapLog == nil {
            log.Fatal(my_errors.ErrorsConfigInitFail + ",请解决键名重复问题,相关键：" + key)
        } else {
            // 程序启动初始化完成
            variable.ZapLog.Warn(my_errors.ErrorsContainerKeyAlreadyExists + ", 相关键：" + key)
        }
    }
    return
}

// Delete 2.删除
func (c *containers) Delete(key string) {
    sMap.Delete(key)
}

// Get 3.传递键，从容器中获取键
func (c *containers) Get(key string) interface{} {
    if value, exists := c.KeyIsExists(key); exists {
        return value
    }
    return nil
}

// KeyIsExists 4. 判断键是否被注册
func (c *containers) KeyIsExists(key string) (interface{}, bool) {
    return sMap.Load(key)
}

// FuzzyDelete 5.按照键的前缀模糊删除容器中注册的内容
func (c *containers) FuzzyDelete(keyPre string) {
    sMap.Range(func(key, value interface{}) bool {
        if keyName, ok := key.(string); ok {
            if strings.HasPrefix(keyName, keyPre) {
                sMap.Delete(keyName)
            }
        }
        return true
    })
}
