package entities

import (
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/GregoryKogan/mephi-databases/internal/models"
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

	if len(s.boardIDs) == 0 || len(s.userIDs) == 0 || len(s.boardRoleIDs) == 0 {
		panic("boardIDs, userIDs or boardRoleIDs are not set")
	}

	for created := uint(0); created < count; {
		boardID := s.boardIDs[rand.Intn(len(s.boardIDs))]
		userID := s.userIDs[rand.Intn(len(s.userIDs))]

		// Check if user is already a member of the board
		var count int64
		s.db.Model(&models.BoardMember{}).Where("board_id = ? AND user_id = ?", boardID, userID).Count(&count)

		if count == 0 {
			roleID := s.boardRoleIDs[rand.Intn(len(s.boardRoleIDs))]

			member := models.BoardMember{
				BoardID:     boardID,
				UserID:      userID,
				BoardRoleID: roleID,
			}

			s.db.Create(&member)
			created++
		}
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
