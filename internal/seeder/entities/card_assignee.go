package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"gorm.io/gorm"
)

type CardAssigneeSeeder interface {
	Seed(count uint)
}

type CardAssigneeSeederImpl struct {
	db *gorm.DB
}

func NewCardAssigneeSeeder(db *gorm.DB) CardAssigneeSeeder {
	return &CardAssigneeSeederImpl{db: db}
}

func (s *CardAssigneeSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d card assignees", count))

	for created := uint(0); created < count; {
		// pick a random card
		card := models.Card{}
		if err := s.db.Model(&models.Card{}).Order("RANDOM()").First(&card).Error; err != nil {
			continue
		}

		list := models.List{}
		if err := s.db.First(&list, card.ListID).Error; err != nil {
			continue
		}

		// pick a random user from same board
		boardMember := models.BoardMember{}
		if err := s.db.Model(&models.BoardMember{}).Where("board_id = ?", list.BoardID).Order("RANDOM()").First(&boardMember).Error; err != nil {
			continue
		}

		// check if user is already assigned to card
		alreadyAssigned := false
		for _, u := range card.Assignees {
			if u.ID == boardMember.UserID {
				alreadyAssigned = true
				break
			}
		}
		if alreadyAssigned {
			continue
		}

		// assign user to card
		var user models.User
		if err := s.db.First(&user, boardMember.UserID).Error; err != nil {
			continue
		}

		if err := s.db.Model(&card).Association("Assignees").Append(&user); err != nil {
			continue
		}

		created++
	}
}
