package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go-demo/config"
)

var (
	LocalConfig *config.LocalConfig = &config.LocalConfig{}
	Db          *gorm.DB
	Redis       *redis.Client
)
