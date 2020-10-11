package app

import (
	"github.com/ghifar/bookstore-oauth-api/domain/access_token"
	"github.com/ghifar/bookstore-oauth-api/domain/db_repository/db"
	"github.com/ghifar/bookstore-oauth-api/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	tokenService := access_token.NewService(db.NewRepository())
	tokenHandler := http.NewHandler(tokenService)

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.POST("/oauth/access_token", tokenHandler.Create)
	router.Run(":8080")
}
