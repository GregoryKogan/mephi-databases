package entities

import (
	"fmt"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type BooksGenresSeeder interface {
	Seed(bookIDs []uint, genreIDs []uint, count uint)
}

type M2MBooksGenres struct {
	BId uint `gorm:"primaryKey;column:b_id"`
	GId uint `gorm:"primaryKey;column:g_id"`
}

type BooksGenresSeederImpl struct {
	db *gorm.DB
}

func NewBooksGenresSeeder(db *gorm.DB) BooksGenresSeeder {
	return &BooksGenresSeederImpl{db: db}
}

func (s *BooksGenresSeederImpl) Seed(bookIDs []uint, genreIDs []uint, count uint) {
	slog.Info(fmt.Sprintf("Seeding %d BookGenres", count))
	defer slog.Info("BookGenres seeded")

	booksGenres := make([]M2MBooksGenres, 0, count)
	uniquePairs := make(map[string]bool)

	created := uint(0)
	for i := 0; i < len(bookIDs) && created < count; i++ {
		booksGenres = append(booksGenres, M2MBooksGenres{
			BId: bookIDs[i],
			GId: genreIDs[i%len(genreIDs)],
		})
		created++

		key := fmt.Sprintf("%d_%d", bookIDs[i], genreIDs[i%len(genreIDs)])
		uniquePairs[key] = true
	}

	for created < count {
		a := gofakeit.Number(0, len(genreIDs)-1)
		b := gofakeit.Number(0, len(bookIDs)-1)
		pair := M2MBooksGenres{
			BId: bookIDs[b],
			GId: genreIDs[a],
		}
		key := fmt.Sprintf("%d_%d", pair.BId, pair.GId)
		if !uniquePairs[key] {
			uniquePairs[key] = true
			booksGenres = append(booksGenres, pair)
			created++
		}
	}

	if err := s.db.Create(&booksGenres).Error; err != nil {
		panic(err)
	}
}

func (M2MBooksGenres) TableName() string {
	return "library.m2m_books_genres"
}
