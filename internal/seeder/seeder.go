package seeder

import (
	"log/slog"

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

	s.roleSeeder.Seed()
	s.boardRoleSeeder.Seed()

	userCount := viper.GetFloat64("seeder.users")
	s.userSeeder.SetRoleIDs(s.roleSeeder.GetIDs())
	s.userSeeder.Seed(uint(userCount))
	s.passwordSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.passwordSeeder.Seed()

	boardCount := userCount * viper.GetFloat64("seeder.boards_per_user")
	s.boardSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardSeeder.Seed(uint(boardCount))

	listCount := boardCount * viper.GetFloat64("seeder.lists_per_board")
	s.listSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.listSeeder.Seed(uint(listCount))

	cardCount := listCount * viper.GetFloat64("seeder.cards_per_list")
	s.cardSeeder.SetListIDs(s.listSeeder.GetIDs())
	s.cardSeeder.Seed(uint(cardCount))

	labelCount := boardCount * viper.GetFloat64("seeder.labels_per_board")
	s.labelSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.labelSeeder.Seed(uint(labelCount))

	commentCount := cardCount * viper.GetFloat64("seeder.comments_per_card")
	s.commentSeeder.SetCardIDs(s.cardSeeder.GetIDs())
	s.commentSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.commentSeeder.Seed(uint(commentCount))

	boardMemberCount := boardCount * viper.GetFloat64("seeder.board_members_per_board")
	s.boardMemberSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.boardMemberSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardMemberSeeder.SetBoardRoleIDs(s.boardRoleSeeder.GetIDs())
	s.boardMemberSeeder.Seed(uint(boardMemberCount))

	attachmentCount := cardCount * viper.GetFloat64("seeder.attachments_per_card")
	s.attachmentSeeder.SetCardIDs(s.cardSeeder.GetIDs())
	s.attachmentSeeder.Seed(uint(attachmentCount))

	cardLabelCount := cardCount * viper.GetFloat64("seeder.card_labels_per_card")
	s.cardLabelSeeder.Seed(uint(cardLabelCount))

	cardAssigneeCount := cardCount * viper.GetFloat64("seeder.card_assignees_per_card")
	s.cardAssigneeSeeder.Seed(uint(cardAssigneeCount))

	slog.Info("Seeding complete")
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
