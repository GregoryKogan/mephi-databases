package seeder

import (
	"log/slog"
	"os"

	"gorm.io/gorm"
)

type Seeder interface {
	Seed()
}

type SeederImpl struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{db: db}
}

func (s *SeederImpl) Seed() {
	s.dropAll()
	s.runInserts()
}

func (s *SeederImpl) dropAll() {
	slog.Info("Dropping old tables")
	tables, err := s.db.Migrator().GetTables()
	if err != nil {
		slog.Error("failed to get tables", slog.Any("error", err))
		panic(err)
	}
	for _, table := range tables {
		if err := s.db.Migrator().DropTable(table); err != nil {
			slog.Error("failed to drop table", slog.Any("table", table), slog.Any("error", err))
			panic(err)
		}
	}
	slog.Info("Tables dropped", slog.Any("tables", tables))
}

func (s *SeederImpl) runInserts() {
	sql, err := os.ReadFile("internal/hw/inserts.sql")
	if err != nil {
		slog.Error("failed to read SQL file", slog.Any("error", err))
		panic(err)
	}

	if err := s.db.Exec(string(sql)).Error; err != nil {
		slog.Error("failed to execute SQL file", slog.Any("error", err))
		panic(err)
	}
}
