package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type BoardSeeder interface {
	Seed(count uint)
	GetIDs() []uint
	SetUserIDs([]uint)
}

type BoardSeederImpl struct {
	db      *gorm.DB
	ids     []uint
	userIDs []uint
}

func NewBoardSeeder(db *gorm.DB) BoardSeeder {
	return &BoardSeederImpl{db: db}
}

func (s *BoardSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d boards", count))
	defer slog.Info("Boards seeded")

	boards := make([]models.Board, count)
	for i := uint(0); i < count; i++ {
		boards[i] = models.Board{
			OwnerID:     s.userIDs[rand.Intn(len(s.userIDs))],
			Title:       "Board " + faker.Word(),
			Description: faker.Sentence(),
		}
	}

	if err := s.db.Create(&boards).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, board := range boards {
		s.ids[i] = board.ID
	}
}

func (s *BoardSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *BoardSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
