package access_token

import (
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
