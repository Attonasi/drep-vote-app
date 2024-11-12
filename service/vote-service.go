package service

import "example/hello/entity"

type VoteService interface {
	Save(entity.Vote) entity.Vote
	FindAll() []entity.Vote
}

type voteService struct {
	votes []entity.Vote
}

func New() VoteService {
	return &voteService{}
}

func (service *voteService) Save(vote entity.Vote) entity.Vote {
	service.votes = append(service.votes, vote)
	return vote
}

func (service *voteService) FindAll() []entity.Vote {
	return service.votes
}
