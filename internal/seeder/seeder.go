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

	// Add count fields
	roleCount         int
	boardRoleCount    int
	userCount         int
	passwordCount     int
	boardCount        int
	labelCount        int
	boardMemberCount  int
	listCount         int
	cardCount         int
	commentCount      int
	attachmentCount   int
	cardLabelCount    int
	cardAssigneeCount int
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

	// Log total records created
	totalRecords := s.calculateTotalRecords()
	slog.Info("Seeding finished",
		slog.Duration("duration", time.Since(startTime)),
		slog.Int("total_records_created", totalRecords),
	)
}

func (s *SeederImpl) seedRoles() {
	s.roleSeeder.Seed()
	s.roleCount = 2 // Assuming there are 2 roles
	s.boardRoleSeeder.Seed()
	s.boardRoleCount = 3 // Assuming there are 3 board roles
}

func (s *SeederImpl) seedUsers() {
	userCount := uint(viper.GetFloat64("seeder.entities.users"))
	s.userSeeder.SetRoleIDs(s.roleSeeder.GetIDs())
	s.userSeeder.Seed(userCount)
	s.userCount = int(userCount)
}

func (s *SeederImpl) seedPasswords() {
	s.passwordSeeder.SetUserRecords(s.userSeeder.GetRecords())
	s.passwordSeeder.Seed()
	s.passwordCount = s.userCount // One password per user
}

func (s *SeederImpl) seedBoards() {
	boardCount := uint(viper.GetFloat64("seeder.entities.boards_per_user") * float64(s.userCount))
	s.boardSeeder.SetUserRecords(s.userSeeder.GetRecords())
	s.boardSeeder.Seed(boardCount)
	s.boardCount = int(boardCount)
}

func (s *SeederImpl) seedLabels() {
	labelCount := uint(viper.GetFloat64("seeder.entities.labels_per_board") * float64(s.boardCount))
	s.labelSeeder.SetBoardRecords(s.boardSeeder.GetRecords())
	s.labelSeeder.Seed(labelCount)
	s.labelCount = int(labelCount)
}

func (s *SeederImpl) seedBoardMembers() {
	boardMemberCount := uint(viper.GetFloat64("seeder.entities.board_members_per_board") * float64(s.boardCount))
	s.boardMemberSeeder.SetBoardRecords(s.boardSeeder.GetRecords())
	s.boardMemberSeeder.SetUserRecords(s.userSeeder.GetRecords())
	s.boardMemberSeeder.SetBoardRoleIDs(s.boardRoleSeeder.GetIDs())
	s.boardMemberSeeder.Seed(boardMemberCount)
	s.boardMemberCount = int(boardMemberCount)
}

func (s *SeederImpl) seedLists() {
	listCount := uint(viper.GetFloat64("seeder.entities.lists_per_board") * float64(s.boardCount))
	s.listSeeder.SetBoardRecords(s.boardSeeder.GetRecords())
	s.listSeeder.Seed(listCount)
	s.listCount = int(listCount)
}

func (s *SeederImpl) seedCards() {
	cardCount := uint(viper.GetFloat64("seeder.entities.cards_per_list") * float64(s.listCount))
	s.cardSeeder.SetListRecords(s.listSeeder.GetRecords())
	s.cardSeeder.Seed(cardCount)
	s.cardCount = int(cardCount)
}

func (s *SeederImpl) seedComments() {
	commentCount := uint(viper.GetFloat64("seeder.entities.comments_per_card") * float64(s.cardCount))
	s.commentSeeder.SetCardRecords(s.cardSeeder.GetRecords())
	s.commentSeeder.SetUserRecords(s.userSeeder.GetRecords())
	s.commentSeeder.Seed(commentCount)
	s.commentCount = int(commentCount)
}

func (s *SeederImpl) seedAttachments() {
	attachmentCount := uint(viper.GetFloat64("seeder.entities.attachments_per_card") * float64(s.cardCount))
	s.attachmentSeeder.SetCardRecords(s.cardSeeder.GetRecords())
	s.attachmentSeeder.Seed(attachmentCount)
	s.attachmentCount = int(attachmentCount)
}

func (s *SeederImpl) seedCardRelations() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cardLabelCount := uint(viper.GetFloat64("seeder.entities.card_labels_per_card") * float64(s.cardCount))
		s.cardLabelSeeder.Seed(cardLabelCount)
		s.cardLabelCount = int(cardLabelCount)
	}()

	go func() {
		defer wg.Done()
		cardAssigneeCount := uint(viper.GetFloat64("seeder.entities.card_assignees_per_card") * float64(s.cardCount))
		s.cardAssigneeSeeder.Seed(cardAssigneeCount)
		s.cardAssigneeCount = int(cardAssigneeCount)
	}()

	wg.Wait()
}

func (s *SeederImpl) calculateTotalRecords() int {
	total := 0
	total += s.roleCount
	total += s.boardRoleCount
	total += s.userCount
	total += s.passwordCount
	total += s.boardCount
	total += s.labelCount
	total += s.boardMemberCount
	total += s.listCount
	total += s.cardCount
	total += s.commentCount
	total += s.attachmentCount
	total += s.cardLabelCount
	total += s.cardAssigneeCount
	return total
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
