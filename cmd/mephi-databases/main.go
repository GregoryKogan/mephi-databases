package main

import (
	"log/slog"
	"os"

	"github.com/GregoryKogan/mephi-databases/internal/config"
	"github.com/GregoryKogan/mephi-databases/internal/seeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.Init()

	// Connect to database
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DSN")}), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect database", slog.Any("error", err))
		panic(err)
	}

	// Seed database
	s := seeder.NewSeeder(db)
	s.Seed()
}
