package db

import (
	"time"

	config "github.com/eminmuhammadi/memcache/config"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// It is a function that opens the database connection.
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

// It is a function that closes the database connection.
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
