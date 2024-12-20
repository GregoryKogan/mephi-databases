package entities

import (
	"fmt"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type SubscriberSeeder interface {
	Seed(count uint)
	GetIDs() []uint
}

type Subscriber struct {
	SId   uint   `gorm:"primaryKey;column:s_id"`
	SName string `gorm:"column:s_name"`
}

type SubscriberSeederImpl struct {
	db  *gorm.DB
	IDs []uint
}

func NewSubscriberSeeder(db *gorm.DB) SubscriberSeeder {
	return &SubscriberSeederImpl{db: db}
}

func (s *SubscriberSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d Subscribers", count))
	defer slog.Info("Subscribers seeded")

	Subscribers := make([]Subscriber, count)
	for i := uint(0); i < count; i++ {
		Subscriber := Subscriber{
			SName: gofakeit.Username(),
		}
		Subscribers[i] = Subscriber
	}

	if err := s.db.Create(&Subscribers).Error; err != nil {
		panic(err)
	}

	s.IDs = make([]uint, count)
	for i, Subscriber := range Subscribers {
		s.IDs[i] = Subscriber.SId
	}
}

func (s *SubscriberSeederImpl) GetIDs() []uint {
	return s.IDs
}
