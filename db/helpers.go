package db

import (
	"strings"
	"time"

	config "github.com/eminmuhammadi/memcache/config"
	"gorm.io/gorm/logger"
)

func TimeNow() time.Time {
	loc, _ := time.LoadLocation(config.DefaultConfig.Timezone)
	return time.Now().In(loc)
}

func TimeNowString() string {
	return TimeNow().Format(time.RFC3339)
}

func GetLogLevel() logger.LogLevel {
	level := strings.ToUpper(config.DefaultConfig.LogLevel)

	switch level {
	case "INFO":
		return logger.Info
	case "WARN":
		return logger.Warn
	case "ERROR":
		return logger.Error
	case "SILENT":
		return logger.Silent
	default:
		return logger.Error
	}
}
