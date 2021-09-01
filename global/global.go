package global

import (
	"github.com/jinzhu/gorm"
	"go-demo/config"
)

var (
	LocalConfig *config.LocalConfig = &config.LocalConfig{}
	Db          *gorm.DB
)
