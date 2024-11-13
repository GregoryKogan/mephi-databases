package seeder

import (
	"log/slog"
	"sync"
	"time"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/entities"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Seeder interface {
	Seed()
}

type SeederImpl struct {
	db                 *gorm.DB
	roleSeeder         entities.RoleSeeder
	boardRoleSeeder    entities.BoardRoleSeeder
	userSeeder         entities.UserSeeder
	passwordSeeder     entities.PasswordSeeder
	boardSeeder        entities.BoardSeeder
	listSeeder         entities.ListSeeder
	cardSeeder         entities.CardSeeder
	labelSeeder        entities.LabelSeeder
	commentSeeder      entities.CommentSeeder
	boardMemberSeeder  entities.BoardMemberSeeder
	attachmentSeeder   entities.AttachmentSeeder
	cardLabelSeeder    entities.CardLabelSeeder
	cardAssigneeSeeder entities.CardAssigneeSeeder
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{
		db:                 db,
		roleSeeder:         entities.NewRoleSeeder(db),
		boardRoleSeeder:    entities.NewBoardRoleSeeder(db),
		userSeeder:         entities.NewUserSeeder(db),
		passwordSeeder:     entities.NewPasswordSeeder(db),
		boardSeeder:        entities.NewBoardSeeder(db),
		listSeeder:         entities.NewListSeeder(db),
		cardSeeder:         entities.NewCardSeeder(db),
		labelSeeder:        entities.NewLabelSeeder(db),
		commentSeeder:      entities.NewCommentSeeder(db),
		boardMemberSeeder:  entities.NewBoardMemberSeeder(db),
		attachmentSeeder:   entities.NewAttachmentSeeder(db),
		cardLabelSeeder:    entities.NewCardLabelSeeder(db),
		cardAssigneeSeeder: entities.NewCardAssigneeSeeder(db),
	}
}

func (s *SeederImpl) Seed() {
	s.prepare()

	slog.Info("Seeding started",
		slog.Int("load_batch_size", viper.GetInt("seeder.load_batch_size")),
		slog.Int("create_batch_size", viper.GetInt("seeder.create_batch_size")),
	)
	startTime := time.Now()
	defer func() {
		slog.Info("Seeding finished", slog.Duration("duration", time.Since(startTime)))
	}()

	// Seed roles and users
	s.seedRoles()
	s.seedUsers()

	// Seed passwords and boards concurrently
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		s.seedPasswords()
	}()
	go func() {
		defer wg.Done()
		s.seedBoards()
	}()
	wg.Wait()

	// Seed labels, board members, and lists concurrently
	wg.Add(3)
	go func() {
		defer wg.Done()
		s.seedLabels()
	}()
	go func() {
		defer wg.Done()
		s.seedBoardMembers()
	}()
	go func() {
		defer wg.Done()
		s.seedLists()
	}()
	wg.Wait()

	// Seed cards
	s.seedCards()

	// Seed comments, attachments, and card relations concurrently
	wg.Add(3)
	go func() {
		defer wg.Done()
		s.seedComments()
	}()
	go func() {
		defer wg.Done()
		s.seedAttachments()
	}()
	go func() {
		defer wg.Done()
		s.seedCardRelations()
	}()
	wg.Wait()
}

func (s *SeederImpl) seedRoles() {
	s.roleSeeder.Seed()
	s.boardRoleSeeder.Seed()
}

func (s *SeederImpl) seedUsers() {
	userCount := viper.GetFloat64("seeder.entities.users")
	s.userSeeder.SetRoleIDs(s.roleSeeder.GetIDs())
	s.userSeeder.Seed(uint(userCount))
}

func (s *SeederImpl) seedPasswords() {
	s.passwordSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.passwordSeeder.Seed()
}

func (s *SeederImpl) seedBoards() {
	boardCount := viper.GetFloat64("seeder.entities.boards_per_user") * viper.GetFloat64("seeder.entities.users")
	s.boardSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardSeeder.Seed(uint(boardCount))
}

func (s *SeederImpl) seedLabels() {
	labelCount := viper.GetFloat64("seeder.entities.labels_per_board") * s.boardSeeder.GetCount()
	s.labelSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.labelSeeder.Seed(uint(labelCount))
}

func (s *SeederImpl) seedBoardMembers() {
	boardMemberCount := viper.GetFloat64("seeder.entities.board_members_per_board") * s.boardSeeder.GetCount()
	s.boardMemberSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.boardMemberSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardMemberSeeder.SetBoardRoleIDs(s.boardRoleSeeder.GetIDs())
	s.boardMemberSeeder.Seed(uint(boardMemberCount))
}

func (s *SeederImpl) seedLists() {
	listCount := viper.GetFloat64("seeder.entities.lists_per_board") * s.boardSeeder.GetCount()
	s.listSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.listSeeder.Seed(uint(listCount))
}

func (s *SeederImpl) seedCards() {
	cardCount := viper.GetFloat64("seeder.entities.cards_per_list") * s.listSeeder.GetCount()
	s.cardSeeder.SetListIDs(s.listSeeder.GetIDs())
	s.cardSeeder.Seed(uint(cardCount))
}

func (s *SeederImpl) seedComments() {
	commentCount := viper.GetFloat64("seeder.entities.comments_per_card") * s.cardSeeder.GetCount()
	s.commentSeeder.SetCardIDs(s.cardSeeder.GetIDs())
	s.commentSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.commentSeeder.Seed(uint(commentCount))
}

func (s *SeederImpl) seedAttachments() {
	attachmentCount := viper.GetFloat64("seeder.entities.attachments_per_card") * s.cardSeeder.GetCount()
	s.attachmentSeeder.SetCardIDs(s.cardSeeder.GetIDs())
	s.attachmentSeeder.Seed(uint(attachmentCount))
}

func (s *SeederImpl) seedCardRelations() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cardLabelCount := viper.GetFloat64("seeder.entities.card_labels_per_card") * s.cardSeeder.GetCount()
		s.cardLabelSeeder.Seed(uint(cardLabelCount))
	}()

	go func() {
		defer wg.Done()
		cardAssigneeCount := viper.GetFloat64("seeder.entities.card_assignees_per_card") * s.cardSeeder.GetCount()
		s.cardAssigneeSeeder.Seed(uint(cardAssigneeCount))
	}()

	wg.Wait()
}

func (s *SeederImpl) prepare() {
	slog.Info("Preparing database")
	s.dropAll()
	s.migrateAll()
}

func (s *SeederImpl) dropAll() {
	slog.Info("Dropping old tables")
	tables, err := s.db.Migrator().GetTables()
	if err != nil {
		slog.Error("failed to get tables", slog.Any("error", err))
		panic(err)
	}
	for _, table := range tables {
		if err := s.db.Migrator().DropTable(table); err != nil {
			slog.Error("failed to drop table", slog.Any("table", table), slog.Any("error", err))
			panic(err)
		}
	}
	slog.Info("Tables dropped", slog.Any("tables", tables))
}

func (s *SeederImpl) migrateAll() {
	slog.Info("Migrating models")
	if err := s.db.AutoMigrate(
		&models.User{},
		&models.Password{},
		&models.Role{},
		&models.Board{},
		&models.BoardMember{},
		&models.BoardRole{},
		&models.List{},
		&models.Card{},
		&models.Label{},
		&models.Comment{},
		&models.Attachment{},
	); err != nil {
		slog.Error("failed to migrate models", slog.Any("error", err))
		panic(err)
	}
}
