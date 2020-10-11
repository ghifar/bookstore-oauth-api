package http

import (
	"github.com/ghifar/bookstore-oauth-api/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandlerInterface interface {
	GetById(ctx *gin.Context)
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
