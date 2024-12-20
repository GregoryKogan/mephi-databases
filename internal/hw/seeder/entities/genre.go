package entities

import (
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type GenreSeeder interface {
	Seed()
	GetIDs() []uint
}

type Genre struct {
	GId   uint   `gorm:"primaryKey;column:g_id"`
	GName string `gorm:"column:g_name"`
}

type GenreSeederImpl struct {
	db  *gorm.DB
	IDs []uint
}

func NewGenreSeeder(db *gorm.DB) GenreSeeder {
	return &GenreSeederImpl{db: db}
}

func (s *GenreSeederImpl) Seed() {
	names := []string{
		"High Fantasy",
		"Urban Fantasy",
		"Grimdark Fantasy",
		"Steampunk",
		"Cyberpunk",
		"Dieselpunk",
		"Gaslamp Fantasy",
		"Science Fiction",
		"Hard Science Fiction",
		"Soft Science Fiction",
		"Space Opera",
		"Military Science Fiction",
		"Post-Apocalyptic",
		"Dystopian",
		"Utopian",
		"Time Travel",
		"Alternate History",
		"Historical Fiction",
		"Historical Romance",
		"Contemporary Romance",
		"Romantic Comedy",
		"Romantic Suspense",
		"Gothic Romance",
		"Paranormal Romance",
		"Cozy Mystery",
		"Crime Thriller",
		"Police Procedural",
		"Legal Thriller",
		"Spy Thriller",
		"Psychological Thriller",
		"Political Thriller",
		"Noir",
		"Supernatural Thriller",
		"Horror",
		"Splatterpunk Horror",
		"Gothic Horror",
		"Creature Horror",
		"Haunted House Horror",
		"Cosmic Horror",
		"Satire",
		"Dark Comedy",
		"Black Humor",
		"Slapstick Comedy",
		"Coming-of-Age",
		"Bildungsroman",
		"Family Saga",
		"Literary Fiction",
		"Experimental Fiction",
		"Stream-of-Consciousness Fiction",
		"Epistolary Fiction",
		"Cli-Fi (Climate Fiction)",
		"Eco-Fiction",
		"Magical Realism",
		"Speculative Fiction",
		"Mythopoeia",
		"Folklore Retellings",
		"Fairy Tale Retellings",
		"Religious Fiction",
		"Spiritual Fiction",
		"Philosophical Fiction",
		"Action Adventure",
		"Survival Fiction",
		"Exploration Fiction",
		"Sea Stories",
		"Western",
		"Weird West",
		"Science Fantasy",
		"Urban Legends",
		"Martial Arts Fiction",
		"Superhero Fiction",
		"Antihero Fiction",
		"Villain Protagonist Fiction",
		"Detective Fiction",
		"Amateur Sleuth Fiction",
		"Techno-Thriller",
		"Medical Thriller",
		"Biopunk",
		"Genetic Engineering Fiction",
		"Hard Boiled Mystery",
		"Cozy Horror",
		"Psychological Horror",
		"Occult Fiction",
		"Zombie Fiction",
		"Vampires and Werewolves",
		"Alien Invasion",
		"First Contact",
		"Space Exploration",
		"AI Uprising",
		"Virtual Reality Fiction",
		"LitRPG",
		"GameLit",
		"Portal Fantasy",
		"Low Fantasy",
		"Sword and Sorcery",
		"Hero's Journey",
		"Moralistic Tales",
		"Postmodern Fiction",
		"Metafiction",
		"Regional Fiction",
		"Tragic Realism",
	}

	count := uint(len(names))

	slog.Info(fmt.Sprintf("Seeding %d Genres", count))
	defer slog.Info("Genres seeded")

	Genres := make([]Genre, count)
	for i := uint(0); i < count; i++ {
		Genre := Genre{
			GName: names[i],
		}
		Genres[i] = Genre
	}

	if err := s.db.Create(&Genres).Error; err != nil {
		panic(err)
	}

	s.IDs = make([]uint, count)
	for i, Genre := range Genres {
		s.IDs[i] = Genre.GId
	}
}

func (s *GenreSeederImpl) GetIDs() []uint {
	return s.IDs
}
