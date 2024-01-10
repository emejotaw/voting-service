package service

import (
	"context"

	"github.com/emejotaw/voting-service/internal/entity"
	"github.com/emejotaw/voting-service/internal/pb"
	"github.com/emejotaw/voting-service/internal/repository"
	"github.com/emejotaw/voting-service/internal/repository/sqlite"
	"github.com/google/uuid"
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

func (s *VoteService) ComputeVote(ctx context.Context, vote *pb.Vote) (*pb.VoteResponse, error) {

	v := &entity.Vote{
		ID:        uuid.New().String(),
		IpAddress: vote.IpAddress,
		UserAgent: vote.UserAgent,
	}

	err := s.repository.Create(v)

	if err != nil {
		return nil, err
	}

	return &pb.VoteResponse{}, nil
}
