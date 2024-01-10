package sqlite

import "gorm.io/gorm"

type SqliteVoteRepository struct {
	db *gorm.DB
}

func NewSqliteVoteRepository(db *gorm.DB) *SqliteVoteRepository {
	return &SqliteVoteRepository{
		db: db,
	}
}
