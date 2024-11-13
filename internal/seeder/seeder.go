package seeder

import (
	"github.com/GregoryKogan/mephi-databases/internal/seeder/entities"
	"gorm.io/gorm"
)

type Seeder interface {
	Seed()
}

type SeederImpl struct {
	roleSeeder      entities.RoleSeeder
	boardRoleSeeder entities.BoardRoleSeeder
	userSeeder      entities.UserSeeder
	passwordSeeder  entities.PasswordSeeder
	boardSeeder     entities.BoardSeeder
	listSeeder      entities.ListSeeder
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{
		roleSeeder:      entities.NewRoleSeeder(db),
		boardRoleSeeder: entities.NewBoardRoleSeeder(db),
		userSeeder:      entities.NewUserSeeder(db),
		passwordSeeder:  entities.NewPasswordSeeder(db),
		boardSeeder:     entities.NewBoardSeeder(db),
		listSeeder:      entities.NewListSeeder(db),
	}
}

func (s *SeederImpl) Seed() {
	s.roleSeeder.Seed()
	s.boardRoleSeeder.Seed()

	s.userSeeder.SetRoleIDs(s.roleSeeder.GetIDs())
	s.userSeeder.Seed(10)

	s.passwordSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.passwordSeeder.Seed()

	s.boardSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.boardSeeder.Seed(10)

	s.listSeeder.SetBoardIDs(s.boardSeeder.GetIDs())
	s.listSeeder.Seed(10)
}
