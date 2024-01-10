package sqlite

import (
	"github.com/emejotaw/voting-service/internal/entity"
	"gorm.io/gorm"
)

type SqliteVoteRepository struct {
	db *gorm.DB
}

func NewSqliteVoteRepository(db *gorm.DB) *SqliteVoteRepository {
	return &SqliteVoteRepository{
		db: db,
	}
}

func (r *SqliteVoteRepository) Create(vote *entity.Vote) error {

	return r.db.Create(vote).Error
}
