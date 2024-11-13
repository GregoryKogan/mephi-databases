package entities

import (
	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
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
	if len(s.userIDs) == 0 {
		panic("userIDs are not set")
	}

	// Delete all passwords before seeding
	s.db.Exec("DELETE FROM passwords")

	for _, userID := range s.userIDs {
		for {
			password := models.Password{
				UserID:    userID,
				Hash:      randomHash(),
				Salt:      randomSalt(),
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

func randomHash() []byte {
	password := faker.Password()
	return []byte(password)
}

func randomSalt() []byte {
	return []byte(faker.Timestamp())
}
