package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/models"
	"github.com/GregoryKogan/mephi-databases/internal/seeder/selector"
	"gorm.io/gorm"
)

type BoardMemberSeeder interface {
	Seed(count uint)
	SetBoardIDs(boardIDs []uint)
	SetUserIDs(userIDs []uint)
	SetBoardRoleIDs(boardRoleIDs []uint)
}

type BoardMemberSeederImpl struct {
	db           *gorm.DB
	boardIDs     []uint
	userIDs      []uint
	boardRoleIDs []uint
}

func NewBoardMemberSeeder(db *gorm.DB) BoardMemberSeeder {
	return &BoardMemberSeederImpl{db: db}
}

func (s *BoardMemberSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d board members", count))
	defer slog.Info("Board members seeded")

	if len(s.boardIDs) == 0 || len(s.userIDs) == 0 || len(s.boardRoleIDs) == 0 {
		panic("boardIDs, userIDs or boardRoleIDs are not set")
	}

	boardMembers := make([]models.BoardMember, count)
	existingMembers := make(map[uint]map[uint]bool)

	for created := uint(0); created < count; {
		boardID := selector.NewSelector().RandomSelect(s.boardIDs)
		userID := selector.NewSelector().RandomSelect(s.userIDs)

		if _, exists := existingMembers[boardID]; !exists {
			existingMembers[boardID] = make(map[uint]bool)
		}

		if !existingMembers[boardID][userID] {
			roleID := selector.NewSelector().ExponentialSelect(s.boardRoleIDs)
			boardMembers[created] = models.BoardMember{
				BoardID:     boardID,
				UserID:      userID,
				BoardRoleID: roleID,
			}
			existingMembers[boardID][userID] = true
			created++
		}
	}

	if err := s.db.Create(&boardMembers).Error; err != nil {
		panic(err)
	}
}

func (s *BoardMemberSeederImpl) SetBoardIDs(ids []uint) {
	s.boardIDs = ids
}

func (s *BoardMemberSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}

func (s *BoardMemberSeederImpl) SetBoardRoleIDs(ids []uint) {
	s.boardRoleIDs = ids
}
