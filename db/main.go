package db

import (
	"time"

	config "github.com/eminmuhammadi/memcache/config"
	sqlite "github.com/glebarez/sqlite"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open(config.DefaultConfig.DBPath),
		&gorm.Config{
			NowFunc: func() time.Time {
				return TimeNow()
			},
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(GetLogLevel()),
		},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
