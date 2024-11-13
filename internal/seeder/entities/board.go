package entities

import (
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
	slog.Info("Seeding boards")

	for i := uint(0); i < count; i++ {
		board := models.Board{
			OwnerID:     s.userIDs[rand.Intn(len(s.userIDs))],
			Title:       "Board " + faker.Word(),
			Description: faker.Sentence(),
		}

		s.db.Create(&board)
		s.ids = append(s.ids, board.ID)
	}

}

func (s *BoardSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *BoardSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
