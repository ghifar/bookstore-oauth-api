package http

import (
	"github.com/ghifar/bookstore-oauth-api/domain/access_token"
	"github.com/ghifar/bookstore-oauth-api/domain/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandlerInterface interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)

}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandlerInterface {
	return &accessTokenHandler{service: service}
}

func (obj *accessTokenHandler) GetById(ctx *gin.Context) {
	tokenId := ctx.Param("access_token_id")
	accessToken, err := obj.service.GetById(tokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}

func (obj *accessTokenHandler) Create(ctx *gin.Context) {
	var accesstoken access_token.AccessToken
	if err := ctx.ShouldBindJSON(&accesstoken); err != nil {
		restErr := errors.NewBadRequestError("invalid jsonBody")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	if err := obj.service.Create(accesstoken); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, accesstoken)

}
