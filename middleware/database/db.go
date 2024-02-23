package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"imserver/config"
)

var DB *gorm.DB

func init() {
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", config.Config.DBConfig.Username, config.Config.DBConfig.Password, config.Config.DBConfig.Host, config.Config.DBConfig.Port, config.Config.DBConfig.DbName, config.Config.DBConfig.Timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("redis connect failed", err.Error())
		panic("failed to connect database" + err.Error())
	}
	DB = db
}
