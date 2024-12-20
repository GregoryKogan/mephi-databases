package entities

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/GregoryKogan/mephi-databases/internal/selector"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type SubscriptionSeeder interface {
	Seed(bookIDs []uint, subscriberIDs []uint, count uint)
	GetIDs() []uint
}

type Subscription struct {
	SbId         uint      `gorm:"primaryKey;column:sb_id"`
	SbSubscriber uint      `gorm:"column:sb_subscriber"`
	SbBook       uint      `gorm:"column:sb_book"`
	SbStart      time.Time `gorm:"column:sb_start"`
	SbFinish     time.Time `gorm:"column:sb_finish"`
	SbIsActive   string    `gorm:"column:sb_is_active"`
}

type SubscriptionSeederImpl struct {
	db  *gorm.DB
	IDs []uint
}

func NewSubscriptionSeeder(db *gorm.DB) SubscriptionSeeder {
	return &SubscriptionSeederImpl{db: db}
}

func (s *SubscriptionSeederImpl) Seed(bookIDs []uint, subscriberIDs []uint, count uint) {
	slog.Info(fmt.Sprintf("Seeding %d Subscriptions", count))
	defer slog.Info("Subscriptions seeded")

	subscriptions := make([]Subscription, 0, count)
	for i := 0; i < int(count); i++ {
		subscriber := gofakeit.Number(0, len(subscriberIDs)-1)
		book := gofakeit.Number(0, len(bookIDs)-1)
		start := selector.NewDateSelector().Before(time.Now(), time.Duration(time.Hour*24*30*12))
		subscription := Subscription{
			SbSubscriber: subscriberIDs[subscriber],
			SbBook:       bookIDs[book],
			SbStart:      start,
		}
		if gofakeit.Bool() {
			subscription.SbFinish = selector.NewDateSelector().Between(start, time.Now())
			subscription.SbIsActive = "N"
		} else {
			subscription.SbIsActive = "Y"
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := s.db.Create(&subscriptions).Error; err != nil {
		panic(err)
	}

	s.IDs = make([]uint, count)
	for i, subscription := range subscriptions {
		s.IDs[i] = subscription.SbId
	}
}

func (s *SubscriptionSeederImpl) GetIDs() []uint {
	return s.IDs
}
