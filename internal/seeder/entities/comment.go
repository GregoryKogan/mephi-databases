package entities

import (
	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
	"golang.org/x/exp/rand"
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
	for i := uint(0); i < count; i++ {
		comment := models.Comment{
			CardID: s.cardIDs[rand.Intn(len(s.cardIDs))],
			UserID: s.userIDs[rand.Intn(len(s.userIDs))],
			Text:   faker.Sentence(),
		}

		s.db.Create(&comment)
	}
}

func (s *CommentSeederImpl) SetCardIDs(ids []uint) {
	s.cardIDs = ids
}

func (s *CommentSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
