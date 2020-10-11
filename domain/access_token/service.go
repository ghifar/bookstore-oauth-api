package access_token

import "github.com/ghifar/bookstore-oauth-api/domain/utils/errors"

type Service interface {
	GetById(id string) (*AccessToken, *errors.RestErr)
}

type Repository interface {
	GetById(id string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetById(id string) (*AccessToken, *errors.RestErr) {
	return s.repository.GetById(id)
}
