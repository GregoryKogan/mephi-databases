package entities

import (
	"fmt"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type BooksAuthorsSeeder interface {
	Seed(bookIDs []uint, authorIDs []uint, count uint)
}

type M2MBooksAuthors struct {
	BId uint `gorm:"primaryKey;column:b_id"`
	AId uint `gorm:"primaryKey;column:a_id"`
}

type BooksAuthorsSeederImpl struct {
	db *gorm.DB
}

func NewBooksAuthorsSeeder(db *gorm.DB) BooksAuthorsSeeder {
	return &BooksAuthorsSeederImpl{db: db}
}

func (s *BooksAuthorsSeederImpl) Seed(bookIDs []uint, authorIDs []uint, count uint) {
	slog.Info(fmt.Sprintf("Seeding %d BookAuthors", count))
	defer slog.Info("BookAuthors seeded")

	booksAuthors := make([]M2MBooksAuthors, 0, count)
	uniquePairs := make(map[string]bool)

	created := uint(0)
	for i := 0; i < len(bookIDs) && created < count; i++ {
		booksAuthors = append(booksAuthors, M2MBooksAuthors{
			BId: bookIDs[i],
			AId: authorIDs[i%len(authorIDs)],
		})
		created++

		key := fmt.Sprintf("%d_%d", bookIDs[i], authorIDs[i%len(authorIDs)])
		uniquePairs[key] = true
	}

	for created < count {
		a := gofakeit.Number(0, len(authorIDs)-1)
		b := gofakeit.Number(0, len(bookIDs)-1)
		pair := M2MBooksAuthors{
			BId: bookIDs[b],
			AId: authorIDs[a],
		}
		key := fmt.Sprintf("%d_%d", pair.BId, pair.AId)
		if !uniquePairs[key] {
			uniquePairs[key] = true
			booksAuthors = append(booksAuthors, pair)
			created++
		}
	}

	if err := s.db.Create(&booksAuthors).Error; err != nil {
		panic(err)
	}
}

func (M2MBooksAuthors) TableName() string {
	return "library.m2m_books_authors"
}
