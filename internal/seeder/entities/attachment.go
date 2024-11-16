package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/brianvoe/gofakeit/v7"
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
	slog.Info(fmt.Sprintf("Seeding %d attachments", count))
	defer slog.Info("Attachments seeded")

	attachments := make([]models.Attachment, count)
	for i := uint(0); i < count; i++ {
		attachments[i] = models.Attachment{
			CardID:  s.cardIDs[rand.Intn(len(s.cardIDs))],
			FileURL: gofakeit.URL(),
		}
	}

	if err := s.db.Create(&attachments).Error; err != nil {
		panic(err)
	}
}

func (s *AttachmentSeederImpl) SetCardIDs(ids []uint) {
	s.cardIDs = ids
}
