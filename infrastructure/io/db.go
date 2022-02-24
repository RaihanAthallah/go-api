package io

import (
	"fmt"
	l "log"
	"os"
	"time"

	"github.com/inventory-management-tokobejo/go-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(isLogged bool) *gorm.DB {
	var dialect gorm.Dialector

	sentryLog := SentryLog{}
	gormConfig := &gorm.Config{}

	if isLogged {
		newLogger := logger.New(
			l.New(os.Stdout, "\r\n", l.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)

		gormConfig.Logger = newLogger
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.DBName,
		config.C.Database.Password,
		config.C.Database.SslMode,
	)
	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		l.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		sentryLog.CaptureError(err)
		l.Fatalln(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	//pool time
	tm := time.Minute * time.Duration(20)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(tm)

	return db
}
