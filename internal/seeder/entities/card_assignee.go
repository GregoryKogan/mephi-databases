package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"golang.org/x/exp/rand"
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
		// pick a random board
		board := models.Board{}
		if err := s.db.Model(&models.Board{}).Order("RANDOM()").Preload("Members").Preload("Lists.Cards.Assignees").First(&board).Error; err != nil {
			continue
		}

		// combine cards from all lists and filter out cards that already have member as assignee
		var cards []models.Card
		for _, list := range board.Lists {
			cards = append(cards, list.Cards...)
		}

		// pick a random card
		if len(cards) == 0 {
			continue
		}
		card := cards[rand.Intn(len(cards))]

		// filter out members that are not already assigned to card
		var members []models.BoardMember
		for _, member := range board.Members {
			alreadyAssigned := false
			for _, assignee := range card.Assignees {
				if member.UserID == assignee.ID {
					alreadyAssigned = true
					break
				}
			}
			if !alreadyAssigned {
				members = append(members, member)
			}
		}

		// pick a random member
		if len(members) == 0 {
			continue
		}
		member := members[rand.Intn(len(members))]

		// assign user to card
		user := models.User{}
		if err := s.db.First(&user, member.UserID).Error; err != nil {
			continue
		}
		card.Assignees = append(card.Assignees, user)

		if err := s.db.Save(&card).Error; err != nil {
			continue
		}

		created++
	}
}
