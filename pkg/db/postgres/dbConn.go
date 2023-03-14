package postgres

import (
	"fmt"
	"log"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDb(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Postgres.PostgresqlHost, c.Postgres.PostgresqlUser, c.Postgres.PostgresqlPassword, c.Postgres.PostgresqlDbname, c.Postgres.PostgresqlPort, c.Postgres.PostgresqlSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connected")
	}
	db.AutoMigrate(&models.User{})
	return db, err
}
