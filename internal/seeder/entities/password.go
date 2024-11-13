package entities

import (
	"crypto/rand"
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
	slog.Info("Seeding passwords")

	if len(s.userIDs) == 0 {
		panic("userIDs are not set")
	}

	for _, userID := range s.userIDs {
		for {
			password := models.Password{
				UserID:    userID,
				Hash:      randomBytes(),
				Salt:      randomBytes(),
				Algorithm: "argon2id",
			}

			if err := s.db.Create(&password).Error; err == nil {
				break
			}
		}
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
