package main

import (
	"fmt"

	"github.com/Sanandrew123/FileSystem/config"
	"github.com/Sanandrew123/FileSystem/internal/api"
)

func main() {
	// loading config
	config.LoadConfig("config/config.yaml")

	// instantiate gin
	r := api.SetupRouter()

	// fmt.Println("服务端口：", config.AppConfig.Server.Port)
	// fmt.Println("存储路径：", config.AppConfig.Storage.Path)

	// run server, listen the port in config
	addr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	r.Run(addr)
}
