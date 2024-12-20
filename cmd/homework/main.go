package main

import (
	"log/slog"
	"os"

	"github.com/GregoryKogan/mephi-databases/internal/config"
	"github.com/GregoryKogan/mephi-databases/internal/hw/seeder"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	config.Init()

	// Connect to database
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DSN")}), &gorm.Config{
		CreateBatchSize: viper.GetInt("seeder.create_batch_size"),
		Logger:          logger.Default.LogMode(logger.Silent),
		// Set schema to library
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "library.",
		},
	})
	if err != nil {
		slog.Error("failed to connect database", slog.Any("error", err))
		panic(err)
	}

	// Seed database
	s := seeder.NewSeeder(db)
	s.Seed()
}
