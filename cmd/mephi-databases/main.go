package main

import (
	"fmt"
	"os"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DSN")}), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database: %v\n", err)
		panic(err)
	}

	// Migrate all models
	if err := db.AutoMigrate(
		&models.User{},
		&models.Password{},
		&models.Role{},
		&models.Board{},
		&models.BoardMember{},
		&models.BoardRole{},
		&models.List{},
		&models.Card{},
		&models.Label{},
		&models.Comment{},
		&models.Attachment{},
	); err != nil {
		fmt.Fprintf(os.Stderr, "failed to migrate database: %v\n", err)
		panic(err)
	}

	// Seed database
	s := seeder.NewSeeder(db)
	s.Seed()
}
