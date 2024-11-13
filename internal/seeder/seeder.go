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
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{
		roleSeeder:      entities.NewRoleSeeder(db),
		boardRoleSeeder: entities.NewBoardRoleSeeder(db),
		userSeeder:      entities.NewUserSeeder(db),
		passwordSeeder:  entities.NewPasswordSeeder(db),
	}
}

func (s *SeederImpl) Seed() {
	s.roleSeeder.Seed()
	s.boardRoleSeeder.Seed()

	s.userSeeder.SetRoleIDs(s.roleSeeder.GetIDs())
	s.userSeeder.Seed(10)

	s.passwordSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.passwordSeeder.Seed()
}
