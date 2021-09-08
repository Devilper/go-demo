package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go-demo/config"
)

var (
	Trans       ut.Translator
	LocalConfig *config.LocalConfig = &config.LocalConfig{}
	Db          *gorm.DB
	Redis       *redis.Client
)
