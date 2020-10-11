package db

import (
	"github.com/ghifar/bookstore-oauth-api/domain/access_token"
	"github.com/ghifar/bookstore-oauth-api/domain/utils/errors"
)

type DbRepository interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Not implemented yet")
}
