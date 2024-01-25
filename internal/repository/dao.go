package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/NikhilSharma03/Mirai/config"
)

type (
	DAO interface {
	}

	dao struct {
		DB *gorm.DB
	}
)

func NewDAO(dbConn *gorm.DB) DAO {
	return &dao{
		DB: dbConn,
	}
}

func ConnectDB(config *config.Config) (*gorm.DB, error) {
	connection := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", config.POSTGRES_HOST, config.POSTGRES_USER, config.POSTGRES_PASSWORD, config.POSTGRES_DB, config.POSTGRES_PORT)
	return gorm.Open(postgres.Open(connection), &gorm.Config{})
}
