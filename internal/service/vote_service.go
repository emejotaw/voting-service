package service

import (
	"github.com/emejotaw/voting-service/internal/pb"
	"github.com/emejotaw/voting-service/internal/repository"
	"github.com/emejotaw/voting-service/internal/repository/sqlite"
	"gorm.io/gorm"
)

type VoteService struct {
	pb.UnimplementedVoteServiceServer
	repository repository.VoteRepository
}

func NewVoteService(db *gorm.DB) *VoteService {
	repository := sqlite.NewSqliteVoteRepository(db)
	return &VoteService{
		repository: repository,
	}
}
