package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/lab/models"
	"github.com/GregoryKogan/mephi-databases/internal/selector"
	"github.com/spf13/viper"
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
	defer slog.Info("Card assignees seeded")

	cardAssignees := make([]models.CardAssignee, 0, count)
	existingAssignees := make(map[uint]map[uint]bool)

	var boards []models.Board
	limit := viper.GetInt("seeder.load_batch_size")
	for created := uint(0); created < count; {
		if err := s.db.Model(&models.Board{}).Order("RANDOM()").Limit(limit).Preload("Members").Preload("Lists.Cards").Find(&boards).Error; err != nil {
			panic(err)
		}

		for _, board := range boards {
			var cards []models.Card
			for _, list := range board.Lists {
				cards = append(cards, list.Cards...)
			}

			if len(cards) == 0 || len(board.Members) == 0 {
				continue
			}

			card := cards[rand.Intn(len(cards))]

			if _, exists := existingAssignees[card.ID]; !exists {
				existingAssignees[card.ID] = make(map[uint]bool)
			}

			for _, member := range board.Members {
				if !existingAssignees[card.ID][member.UserID] {
					minAssignmentDate := card.CreatedAt
					if member.CreatedAt.After(minAssignmentDate) {
						minAssignmentDate = member.CreatedAt
					}

					if minAssignmentDate.After(card.DueDate.Time) {
						continue
					}

					cardAssignees = append(cardAssignees, models.CardAssignee{
						CardID: card.ID,
						UserID: member.UserID,
						Model: gorm.Model{
							CreatedAt: selector.NewDateSelector().Between(minAssignmentDate, card.DueDate.Time),
						},
					})
					existingAssignees[card.ID][member.UserID] = true
					created++
					break
				}
			}

			if created >= count {
				break
			}
		}
	}

	if err := s.db.Create(&cardAssignees).Error; err != nil {
		panic(err)
	}
}
