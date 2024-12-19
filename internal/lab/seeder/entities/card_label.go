package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/lab/models"
	"github.com/spf13/viper"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type CardLabelSeeder interface {
	Seed(count uint)
}

type CardLabelSeederImpl struct {
	db *gorm.DB
}

func NewCardLabelSeeder(db *gorm.DB) CardLabelSeeder {
	return &CardLabelSeederImpl{db: db}
}

func (s *CardLabelSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d card labels", count))
	defer slog.Info("Card labels seeded")

	cardLabels := make([]map[string]interface{}, 0, count)
	existingCardLabels := make(map[uint]map[uint]bool)

	var boards []models.Board
	limit := viper.GetInt("seeder.load_batch_size")
	for created := uint(0); created < count; {
		if err := s.db.Model(&models.Board{}).Order("RANDOM()").Limit(limit).Preload("Labels").Preload("Lists.Cards").Find(&boards).Error; err != nil {
			panic(err)
		}

		for _, board := range boards {
			if len(board.Labels) == 0 {
				continue
			}
			label := board.Labels[rand.Intn(len(board.Labels))]

			var cards []models.Card
			for _, list := range board.Lists {
				cards = append(cards, list.Cards...)
			}

			if len(cards) == 0 {
				continue
			}
			card := cards[rand.Intn(len(cards))]

			if _, exists := existingCardLabels[card.ID]; !exists {
				existingCardLabels[card.ID] = make(map[uint]bool)
			}

			if !existingCardLabels[card.ID][label.ID] {
				cardLabels = append(cardLabels, map[string]interface{}{
					"card_id":  card.ID,
					"label_id": label.ID,
				})
				existingCardLabels[card.ID][label.ID] = true
				created++
			}

			if created >= count {
				break
			}
		}
	}

	if err := s.db.Table("card_labels").Create(&cardLabels).Error; err != nil {
		panic(err)
	}
}
