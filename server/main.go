package main

import (
	"bbs-go/cmd"
	"bbs-go/model"
	"bbs-go/scheduler"
	_ "bbs-go/service/eventhandler"
	"bbs-go/util/common"
	"bbs-go/util/config"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"bbs-go/util/simple"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var configFile = flag.String("config", "./bbs-go.yaml", "配置文件路径")

func init() {
	flag.Parse()

	// 初始化配置
	conf := config.Init(*configFile)

	// gorm配置
	gormConf := &gorm.Config{}

	// 初始化日志
	if file, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		if conf.ShowSql {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
	}

	// 连接数据库
	var url string
	if conf.Database.Driver == "sqlite" {
		url = conf.Database.Sqlite.Path
	} else {
		url = conf.Database.Mysql.Url
	}
	if err := simple.OpenDB(conf.Database.Driver, url, gormConf, 10, 20, model.Models...); err != nil {
		logrus.Error(err)
	}

}

func main() {
	if common.IsProd() {
		// 开启定时任务
		scheduler.Start()
	}

	cmd.Web()
}
