package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(), "brand new access token should not be expired")

	assert.Equal(t, "", at.AccessToken, "token should not have defined access token id")

	assert.True(t, at.UserId == 0, "token shouldnt have an associated user id")
}
