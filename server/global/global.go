package global

import (
	"github.com/sirupsen/logrus"
	"go-blog/config"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MySQLLog logger.Interface
)
