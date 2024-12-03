package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type CommentSeeder interface {
	Seed(count uint)
	SetCardIDs(cardIDs []uint)
	SetUserIDs(userIDs []uint)
}

type CommentSeederImpl struct {
	db      *gorm.DB
	cardIDs []uint
	userIDs []uint
}

func NewCommentSeeder(db *gorm.DB) CommentSeeder {
	return &CommentSeederImpl{db: db}
}

func (s *CommentSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d comments", count))
	defer slog.Info("Comments seeded")

	comments := make([]models.Comment, count)
	for i := uint(0); i < count; i++ {
		comments[i] = models.Comment{
			CardID: selector.NewSliceSelector().Random(s.cardIDs),
			UserID: selector.NewSliceSelector().Random(s.userIDs),
			Text:   gofakeit.Comment(),
		}
	}

	if err := s.db.Create(&comments).Error; err != nil {
		panic(err)
	}
}

func (s *CommentSeederImpl) SetCardIDs(ids []uint) {
	s.cardIDs = ids
}

func (s *CommentSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
