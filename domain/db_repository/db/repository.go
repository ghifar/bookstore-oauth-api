package db

import (
	"fmt"
	"github.com/ghifar/bookstore-oauth-api/clients/cassandra"
	"github.com/ghifar/bookstore-oauth-api/domain/access_token"
	"github.com/ghifar/bookstore-oauth-api/domain/utils/errors"
)

const (
	Q_GET_ACCESS_TOKEN    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token= ?;"
	Q_CREATE_ACCESS_TOKEN = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	Q_UPDATE_ACCESS_TOKEN = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var res access_token.AccessToken
	if err := cassandra.GetSession().Query(Q_GET_ACCESS_TOKEN, id).Scan(
		&res.AccessToken,
		&res.UserId,
		&res.ClientId,
		&res.Expires,
	); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &res, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(Q_CREATE_ACCESS_TOKEN,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save access token in database: %s", err))
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr{
	if err := cassandra.GetSession().Query(Q_UPDATE_ACCESS_TOKEN, token.Expires, token.AccessToken ).Exec() ; err!= nil{
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save access token in database: %s", err))
	}
	return nil
}
