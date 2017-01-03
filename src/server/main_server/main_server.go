package main

//服务主程序

import (
	"frame/logger"
	"time"
)

func main() {
	logger.Info("-----------------------Main server start-----------------------")
	for {
		time.Sleep(1 * time.Second)
	}
}
