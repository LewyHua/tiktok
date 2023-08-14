// Code generated by hertz generator.

package main

import (
	"douyin/config"
	"douyin/logger"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertz_config "github.com/cloudwego/hertz/pkg/common/config"
)

func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 加载日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	h := server.Default(hertz_config.Option{F: func(o *hertz_config.Options) {
		o.Addr = "127.0.0.1:8081"
	}})
	customizedRegister(h)
	h.Spin()
}
