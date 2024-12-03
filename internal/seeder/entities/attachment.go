package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
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
			CardID:  selector.NewSliceSelector().Random(s.cardIDs),
			FileURL: gofakeit.DomainName() + "/" + gofakeit.Word() + "." + gofakeit.FileExtension(),
		}
	}

	if err := s.db.Create(&attachments).Error; err != nil {
		panic(err)
	}
}

func (s *AttachmentSeederImpl) SetCardIDs(ids []uint) {
	s.cardIDs = ids
}
