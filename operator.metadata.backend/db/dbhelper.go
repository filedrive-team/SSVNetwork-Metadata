package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"ssv_operator_metadata/db/models"
	"ssv_operator_metadata/db/repos"
)

func NewDBGorm(dataSourceName string) *gorm.DB {
	var dialector gorm.Dialector
	dbUrl := dataSourceName
	switch "mysql" {
	case "postgres":
		dialector = postgres.Open(dbUrl)
	case "mysql":
		dialector = mysql.Open(dbUrl)
	}
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold: 5 * time.Second, // slow SQL threshold
			LogLevel:      logger.Error,    // Log level
			Colorful:      true,            // disable color printing
		},
	)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("database setup err: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("db.DB() err: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(models.DbModels()...)

	return db
}

type DBManager struct {
	*repos.ReposCli
}

func NewDBManager(dsn string) *DBManager {
	return &DBManager{
		repos.NewRepoCli(NewDBGorm(dsn)),
	}
}
