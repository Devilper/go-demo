package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-demo/global"
	"go-demo/model"
)

func InitDb() {
	driver_name := global.LocalConfig.Db.DriverName
	host := global.LocalConfig.Db.Host
	port := global.LocalConfig.Db.Port
	database := global.LocalConfig.Db.Database
	username := global.LocalConfig.Db.Username
	password := global.LocalConfig.Db.Password
	charset := global.LocalConfig.Db.Charset

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driver_name, args)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&model.User{})
	global.Db = db
}

func InitRedis() error {
	address := global.LocalConfig.Redis.Address
	password := global.LocalConfig.Redis.Password
	db := global.LocalConfig.Redis.Db
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	_, err := global.Redis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	InitDb()
	if err := InitRedis(); err != nil {
		panic(err)
	}
	//defer global.Db.Close()
}
