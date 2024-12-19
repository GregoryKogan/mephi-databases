package entities

import (
	"crypto/rand"
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/lab/models"
	"gorm.io/gorm"
)

type PasswordSeeder interface {
	Seed()
	SetUserRecords([]Record)
}

type PasswordSeederImpl struct {
	db          *gorm.DB
	userRecords []Record
}

func NewPasswordSeeder(db *gorm.DB) PasswordSeeder {
	return &PasswordSeederImpl{db: db}
}

func (s *PasswordSeederImpl) Seed() {
	slog.Info(fmt.Sprintf("Seeding %d passwords", len(s.userRecords)))
	defer slog.Info("Passwords seeded")

	if len(s.userRecords) == 0 {
		panic("userIDs are not set")
	}

	passwords := make([]models.Password, len(s.userRecords))
	for i, record := range s.userRecords {
		passwords[i] = models.Password{
			UserID:    record.ID,
			Hash:      randomBytes(),
			Salt:      randomBytes(),
			Algorithm: "argon2id",
		}
	}

	if err := s.db.Create(&passwords).Error; err != nil {
		panic(err)
	}
}

func (s *PasswordSeederImpl) SetUserRecords(records []Record) {
	s.userRecords = records
}

func randomBytes() []byte {
	bytes := make([]byte, 128)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil
	}
	return bytes
}
