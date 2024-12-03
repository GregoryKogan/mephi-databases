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
	SetCardRecords([]Record)
}

type AttachmentSeederImpl struct {
	db          *gorm.DB
	cardRecords []Record
}

func NewAttachmentSeeder(db *gorm.DB) AttachmentSeeder {
	return &AttachmentSeederImpl{db: db}
}

func (s *AttachmentSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d attachments", count))
	defer slog.Info("Attachments seeded")

	attachments := make([]models.Attachment, count)
	for i := uint(0); i < count; i++ {
		cardRecord := s.cardRecords[selector.NewSliceSelector().Random(len(s.cardRecords))]
		attachments[i] = models.Attachment{
			CardID:  cardRecord.ID,
			FileURL: gofakeit.DomainName() + "/" + gofakeit.Word() + "." + gofakeit.FileExtension(),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().BeforeNow(cardRecord.CreatedAt),
			},
		}
	}

	if err := s.db.Create(&attachments).Error; err != nil {
		panic(err)
	}
}

func (s *AttachmentSeederImpl) SetCardRecords(records []Record) {
	s.cardRecords = records
}
