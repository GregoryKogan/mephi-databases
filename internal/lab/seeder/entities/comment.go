package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/lab/models"
	"github.com/GregoryKogan/mephi-databases/internal/lab/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type CommentSeeder interface {
	Seed(count uint)
	SetCardRecords([]Record)
	SetUserRecords([]Record)
}

type CommentSeederImpl struct {
	db          *gorm.DB
	cardRecords []Record
	userRecords []Record
}

func NewCommentSeeder(db *gorm.DB) CommentSeeder {
	return &CommentSeederImpl{db: db}
}

func (s *CommentSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d comments", count))
	defer slog.Info("Comments seeded")

	comments := make([]models.Comment, count)
	for i := uint(0); i < count; i++ {
		cardRecord := s.cardRecords[selector.NewSliceSelector().Random(len(s.cardRecords))]
		userRecord := s.userRecords[selector.NewSliceSelector().Random(len(s.userRecords))]

		minCommentDate := cardRecord.CreatedAt
		if userRecord.CreatedAt.After(minCommentDate) {
			minCommentDate = userRecord.CreatedAt
		}

		comments[i] = models.Comment{
			CardID: cardRecord.ID,
			UserID: userRecord.ID,
			Text:   gofakeit.Comment(),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().BeforeNow(minCommentDate),
			},
		}
	}

	if err := s.db.Create(&comments).Error; err != nil {
		panic(err)
	}
}

func (s *CommentSeederImpl) SetCardRecords(records []Record) {
	s.cardRecords = records
}

func (s *CommentSeederImpl) SetUserRecords(records []Record) {
	s.userRecords = records
}
