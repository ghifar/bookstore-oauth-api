package access_token

import (
	"github.com/ghifar/bookstore-oauth-api/domain/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		AccessToken: "",
		UserId:      0,
		ClientId:    0,
		Expires:     time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken *AccessToken) isExpired() bool {
	return false
}

func (accessToken *AccessToken) Validate() *errors.RestErr {
	accessTokenId := strings.TrimSpace(accessToken.AccessToken)
	if len(accessTokenId) == 0 {
		return errors.NewBadRequestError("invalid access token")
	}

	if accessToken.UserId <= 0 {
		return errors.NewBadRequestError("invalid access token")
	}
	if accessToken.ClientId <= 0 {
		return errors.NewBadRequestError("invalid clientid")
	}
	if accessToken.Expires <= 0 {
		return errors.NewBadRequestError("invalid expires")
	}
	return nil
}