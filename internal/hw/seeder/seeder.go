package seeder

import (
	"log/slog"
	"os"
	"sync"

	"github.com/GregoryKogan/mephi-databases/internal/hw/seeder/entities"
	"gorm.io/gorm"
)

type Seeder interface {
	Seed()
}

type SeederImpl struct {
	db                *gorm.DB
	authorSeeder      entities.AuthorSeeder
	bookSeeder        entities.BookSeeder
	genreSeeder       entities.GenreSeeder
	subscriberSeeder  entities.SubscriberSeeder
	booksAuthorSeeder entities.BooksAuthorsSeeder
	booksGenreSeeder  entities.BooksGenresSeeder
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{
		db:                db,
		authorSeeder:      entities.NewAuthorSeeder(db),
		bookSeeder:        entities.NewBookSeeder(db),
		genreSeeder:       entities.NewGenreSeeder(db),
		subscriberSeeder:  entities.NewSubscriberSeeder(db),
		booksAuthorSeeder: entities.NewBooksAuthorsSeeder(db),
		booksGenreSeeder:  entities.NewBooksGenresSeeder(db),
	}
}

func (s *SeederImpl) Seed() {
	s.dropAll()
	s.runInserts()

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		s.authorSeeder.Seed(10_000)
	}()
	go func() {
		defer wg.Done()
		s.bookSeeder.Seed(100_000)
	}()
	go func() {
		defer wg.Done()
		s.genreSeeder.Seed()
	}()
	go func() {
		defer wg.Done()
		s.subscriberSeeder.Seed(1_000_000)
	}()
	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		s.booksAuthorSeeder.Seed(s.bookSeeder.GetIDs(), s.authorSeeder.GetIDs(), 1_000_000)
	}()
	go func() {
		defer wg.Done()
		s.booksGenreSeeder.Seed(s.bookSeeder.GetIDs(), s.genreSeeder.GetIDs(), 1_000_000)
	}()
	wg.Wait()
}

func (s *SeederImpl) dropAll() {
	schema := "library"
	tables := []string{"authors", "books", "genres", "m2m_books_authors", "m2m_books_genres", "subscribers", "subscriptions"}
	for _, table := range tables {
		fullTable := schema + "." + table
		if err := s.db.Migrator().DropTable(fullTable); err != nil {
			slog.Error("failed to drop table", slog.Any("table", table), slog.Any("error", err))
			panic(err)
		}
	}
	slog.Info("Tables dropped", slog.Any("tables", tables))
}

func (s *SeederImpl) runInserts() {
	sql, err := os.ReadFile("internal/hw/init.sql")
	if err != nil {
		slog.Error("failed to read SQL file", slog.Any("error", err))
		panic(err)
	}

	if err := s.db.Exec(string(sql)).Error; err != nil {
		slog.Error("failed to execute SQL file", slog.Any("error", err))
		panic(err)
	}
}
