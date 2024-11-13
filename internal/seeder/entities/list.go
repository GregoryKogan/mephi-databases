package entities

import (
	"math/rand"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type ListSeeder interface {
	Seed(count uint)
	GetIDs() []uint
	SetBoardIDs([]uint)
}

type ListSeederImpl struct {
	db       *gorm.DB
	ids      []uint
	boardIDs []uint
}

func NewListSeeder(db *gorm.DB) ListSeeder {
	return &ListSeederImpl{db: db}
}

func (s *ListSeederImpl) Seed(count uint) {
	if len(s.boardIDs) == 0 {
		panic("boardIDs are not set")
	}

	// Delete all lists before seeding
	s.db.Exec("DELETE FROM lists")

	for i := uint(0); i < count; i++ {
		list := models.List{
			BoardID: s.boardIDs[rand.Intn(len(s.boardIDs))],
			Title:   "List " + faker.Word(),
			Order:   int(i),
		}

		s.db.Create(&list)
		s.ids = append(s.ids, list.ID)
	}
}

func (s *ListSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *ListSeederImpl) SetBoardIDs(ids []uint) {
	s.boardIDs = ids
}
