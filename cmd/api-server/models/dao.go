package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DAO struct {
	DB *gorm.DB
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

const (
	defaultSSLModel = "disable"
	defaultTimeZone = "Asia/Taipei"
)

func NewDAO(config DBConfig) (*DAO, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			config.DBHost,
			config.DBUsername,
			config.DBPassword,
			config.DBName,
			config.DBPort,
			defaultSSLModel,
			defaultTimeZone,
		),
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to open db: %v", err)
	}

	return &DAO{
		DB: db,
	}, nil
}

func (d *DAO) Init() error {
	if err := d.DB.AutoMigrate(Product{}); err != nil {
		return fmt.Errorf("failed to migrate the db: %v", err)
	}
	return nil
}
