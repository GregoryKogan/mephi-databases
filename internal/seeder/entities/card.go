package entities

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type CardSeeder interface {
	Seed(count uint)
	GetRecords() []Record
	GetCount() float64
	SetListRecords([]Record)
}

type CardSeederImpl struct {
	db          *gorm.DB
	records     []Record
	listRecords []Record
}

func NewCardSeeder(db *gorm.DB) CardSeeder {
	return &CardSeederImpl{db: db}
}

func (s *CardSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d cards", count))
	defer slog.Info("Cards seeded")

	if len(s.listRecords) == 0 {
		panic("listIDs are not set")
	}

	cards := make([]models.Card, count)
	for i := uint(0); i < count; i++ {
		listRecord := s.listRecords[selector.NewSliceSelector().Random(len(s.listRecords))]
		cards[i] = models.Card{
			ListID:    listRecord.ID,
			Title:     cases.Title(language.English, cases.Compact).String(gofakeit.Adjective() + " " + gofakeit.Noun()),
			Content:   gofakeit.Paragraph(2, 3, 10, " "),
			Completed: selector.NewBoolSelector().WithProbability(0.7),
			DueDate:   sql.NullTime{Time: selector.NewDateSelector().After(time.Now(), time.Duration(time.Hour*24*30)), Valid: true},
			Order:     int(i),
			Model: gorm.Model{
				CreatedAt: selector.NewDateSelector().BeforeNow(listRecord.CreatedAt),
			},
		}
	}

	if err := s.db.Create(&cards).Error; err != nil {
		panic(err)
	}

	s.records = make([]Record, count)
	for i, card := range cards {
		s.records[i] = Record{ID: card.ID, CreatedAt: card.CreatedAt}
	}
}

func (s *CardSeederImpl) GetRecords() []Record {
	return s.records
}

func (s *CardSeederImpl) GetCount() float64 {
	return float64(len(s.records))
}

func (s *CardSeederImpl) SetListRecords(records []Record) {
	s.listRecords = records
}
