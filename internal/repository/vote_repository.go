package repository

import "github.com/emejotaw/voting-service/internal/entity"

type VoteRepository interface {
	Create(vote *entity.Vote) error
}
