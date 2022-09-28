package internal

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"goshop/account_srv/model"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

type DataBase struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	DBName string `mapstructure:"dbName"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func InitDB(){
	// logger配置
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      true,         // 禁用彩色打印
		},
	)

	// 1.连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ViperConf.DataBase.Username,
		ViperConf.DataBase.Password,
		ViperConf.DataBase.Host,
		ViperConf.DataBase.Port,
		ViperConf.DataBase.DBName)
	zap.S().Info(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log})
	if err !=nil{
		panic("Mysql connect faild:"+ err.Error())
	}

	err = DB.AutoMigrate(&model.Account{})
	if err !=nil{
		fmt.Println(err)
	}
}
