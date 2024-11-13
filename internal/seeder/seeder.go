package seeder

import (
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/seeder/entities"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Seeder interface {
	Seed()
}

type SeederImpl struct {
	db                *gorm.DB
	roleSeeder        entities.RoleSeeder
	boardRoleSeeder   entities.BoardRoleSeeder
	userSeeder        entities.UserSeeder
	passwordSeeder    entities.PasswordSeeder
	boardSeeder       entities.BoardSeeder
	listSeeder        entities.ListSeeder
	cardSeeder        entities.CardSeeder
	labelSeeder       entities.LabelSeeder
	commentSeeder     entities.CommentSeeder
	boardMemberSeeder entities.BoardMemberSeeder
	attachmentSeeder  entities.AttachmentSeeder
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{
		db:                db,
		roleSeeder:        entities.NewRoleSeeder(db),
		boardRoleSeeder:   entities.NewBoardRoleSeeder(db),
		userSeeder:        entities.NewUserSeeder(db),
		passwordSeeder:    entities.NewPasswordSeeder(db),
		boardSeeder:       entities.NewBoardSeeder(db),
		listSeeder:        entities.NewListSeeder(db),
		cardSeeder:        entities.NewCardSeeder(db),
		labelSeeder:       entities.NewLabelSeeder(db),
		commentSeeder:     entities.NewCommentSeeder(db),
		boardMemberSeeder: entities.NewBoardMemberSeeder(db),
		attachmentSeeder:  entities.NewAttachmentSeeder(db),
	}
}

func (s *SeederImpl) Seed() {
	s.deleteAll()

	s.roleSeeder.Seed()
	s.boardRoleSeeder.Seed()

	s.userSeeder.SetRoleIDs(s.roleSeeder.GetIDs())
	s.userSeeder.Seed(viper.GetUint("seeder.users"))

	s.passwordSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.passwordSeeder.Seed()

	s.boardSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardSeeder.Seed(viper.GetUint("seeder.boards"))

	s.listSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.listSeeder.Seed(viper.GetUint("seeder.lists"))

	s.cardSeeder.SetListIDs(s.listSeeder.GetIDs())
	s.cardSeeder.Seed(viper.GetUint("seeder.cards"))

	s.labelSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.labelSeeder.Seed(viper.GetUint("seeder.labels"))

	s.commentSeeder.SetCardIDs(s.cardSeeder.GetIDs())
	s.commentSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.commentSeeder.Seed(viper.GetUint("seeder.comments"))

	s.boardMemberSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.boardMemberSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardMemberSeeder.SetBoardRoleIDs(s.boardRoleSeeder.GetIDs())
	s.boardMemberSeeder.Seed(viper.GetUint("seeder.board_members"))

	s.attachmentSeeder.SetCardIDs(s.cardSeeder.GetIDs())
	s.attachmentSeeder.Seed(viper.GetUint("seeder.attachments"))

	slog.Info("Seeding complete")
}

func (s *SeederImpl) deleteAll() {
	slog.Info("Deleting all data")
	s.db.Exec("DELETE FROM passwords")
	s.db.Exec("DELETE FROM attachments")
	s.db.Exec("DELETE FROM comments")
	s.db.Exec("DELETE FROM board_members")
	s.db.Exec("DELETE FROM labels")
	s.db.Exec("DELETE FROM cards")
	s.db.Exec("DELETE FROM lists")
	s.db.Exec("DELETE FROM boards")
	s.db.Exec("DELETE FROM users")
	s.db.Exec("DELETE FROM roles")
	s.db.Exec("DELETE FROM board_roles")
}
