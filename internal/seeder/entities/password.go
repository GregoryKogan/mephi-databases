package entities

import (
	"crypto/rand"
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"gorm.io/gorm"
)

type PasswordSeeder interface {
	Seed()
	SetUserIDs(ids []uint)
}

type PasswordSeederImpl struct {
	db      *gorm.DB
	userIDs []uint
}

func NewPasswordSeeder(db *gorm.DB) PasswordSeeder {
	return &PasswordSeederImpl{db: db}
}

func (s *PasswordSeederImpl) Seed() {
	slog.Info(fmt.Sprintf("Seeding %d passwords", len(s.userIDs)))
	defer slog.Info("Passwords seeded")

	if len(s.userIDs) == 0 {
		panic("userIDs are not set")
	}

	passwords := make([]models.Password, len(s.userIDs))
	for i, userID := range s.userIDs {
		passwords[i] = models.Password{
			UserID:    userID,
			Hash:      randomBytes(),
			Salt:      randomBytes(),
			Algorithm: "argon2id",
		}
	}

	if err := s.db.Create(&passwords).Error; err != nil {
		panic(err)
	}
}

func (s *PasswordSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}

func randomBytes() []byte {
	bytes := make([]byte, 128)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil
	}
	return bytes
}
