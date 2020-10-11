package db

import (
	"github.com/ghifar/bookstore-oauth-api/clients/cassandra"
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
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("Not implemented yet")
}
