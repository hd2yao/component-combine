package variable

import (
    "log"
    "os"
    "strings"

    "go.uber.org/zap"

    "component-combine/app/global/my_errors"
    "component-combine/app/utils/yml_config/interf"
)

var (
    BasePath        string      // 定义项目的根目录
    ConfigKeyPrefix = "Config_" // 配置文件键值缓存时，键的前缀

    ZapLog    *zap.Logger            // 全局日志指针
    ConfigYml interf.YmlConfigInterf // 全局配置文件指针
)

func init() {
    // 1.初始化程序根目录
    if path, err := os.Getwd(); err == nil {
        // 路径进行处理，兼容单元测试程序启动时的奇怪路径
        if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
            BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `\test`, "", 1)
        } else {
            BasePath = path
        }
    } else {
        log.Fatal(my_errors.ErrorsBasePath)
    }
}
