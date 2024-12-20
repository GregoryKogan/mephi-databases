package entities

import (
	"fmt"
	"log/slog"

	"github.com/GregoryKogan/mephi-databases/internal/lab/models"
	"github.com/GregoryKogan/mephi-databases/internal/selector"
	"gorm.io/gorm"
)

type BoardMemberSeeder interface {
	Seed(count uint)
	SetBoardRecords([]Record)
	SetUserRecords([]Record)
	SetBoardRoleIDs([]uint)
}

type BoardMemberSeederImpl struct {
	db           *gorm.DB
	boardRecords []Record
	userRecords  []Record
	boardRoleIDs []uint
}

func NewBoardMemberSeeder(db *gorm.DB) BoardMemberSeeder {
	return &BoardMemberSeederImpl{db: db}
}

func (s *BoardMemberSeederImpl) Seed(count uint) {
	slog.Info(fmt.Sprintf("Seeding %d board members", count))
	defer slog.Info("Board members seeded")

	if len(s.boardRecords) == 0 || len(s.userRecords) == 0 || len(s.boardRoleIDs) == 0 {
		panic("boardIDs, userIDs or boardRoleIDs are not set")
	}

	boardMembers := make([]models.BoardMember, count)
	existingMembers := make(map[uint]map[uint]bool)

	for created := uint(0); created < count; {
		boardRecord := s.boardRecords[selector.NewSliceSelector().Random(len(s.boardRecords))]
		userRecord := s.userRecords[selector.NewSliceSelector().Random(len(s.userRecords))]

		if _, exists := existingMembers[boardRecord.ID]; !exists {
			existingMembers[boardRecord.ID] = make(map[uint]bool)
		}

		if !existingMembers[boardRecord.ID][userRecord.ID] {
			roleID := s.boardRoleIDs[selector.NewSliceSelector().Exponential(len(s.boardRoleIDs))]

			minJoinDate := boardRecord.CreatedAt
			if userRecord.CreatedAt.After(minJoinDate) {
				minJoinDate = userRecord.CreatedAt
			}

			boardMembers[created] = models.BoardMember{
				BoardID:     boardRecord.ID,
				UserID:      userRecord.ID,
				BoardRoleID: roleID,
				Model: gorm.Model{
					CreatedAt: selector.NewDateSelector().BeforeNow(minJoinDate),
				},
			}
			existingMembers[boardRecord.ID][userRecord.ID] = true
			created++
		}
	}

	if err := s.db.Create(&boardMembers).Error; err != nil {
		panic(err)
	}
}

func (s *BoardMemberSeederImpl) SetBoardRecords(records []Record) {
	s.boardRecords = records
}

func (s *BoardMemberSeederImpl) SetUserRecords(records []Record) {
	s.userRecords = records
}

func (s *BoardMemberSeederImpl) SetBoardRoleIDs(ids []uint) {
	s.boardRoleIDs = ids
}
