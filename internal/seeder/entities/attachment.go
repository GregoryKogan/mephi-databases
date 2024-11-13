package entities

import (
	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type AttachmentSeeder interface {
	Seed(count uint)
	SetCardIDs(cardIDs []uint)
}

type AttachmentSeederImpl struct {
	db      *gorm.DB
	cardIDs []uint
}

func NewAttachmentSeeder(db *gorm.DB) AttachmentSeeder {
	return &AttachmentSeederImpl{db: db}
}

func (s *AttachmentSeederImpl) Seed(count uint) {
	for i := uint(0); i < count; i++ {
		attachment := models.Attachment{
			CardID:  s.cardIDs[rand.Intn(len(s.cardIDs))],
			FileURL: faker.URL(),
		}

		s.db.Create(&attachment)
	}
}

func (s *AttachmentSeederImpl) SetCardIDs(ids []uint) {
	s.cardIDs = ids
}
