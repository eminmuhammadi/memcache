package db

import (
	"strings"
	"time"

	config "github.com/eminmuhammadi/memcache/config"
	"gorm.io/gorm/logger"
)

// TimeNow returns the current time in the configured timezone.
func TimeNow() time.Time {
	loc, _ := time.LoadLocation(config.DefaultConfig.Timezone)
	return time.Now().In(loc)
}

// TimeNowString returns the current time in the configured timezone as a string.
func TimeNowString() string {
	return TimeNow().Format(time.RFC3339)
}

// GetLogLevel returns the log level.
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
