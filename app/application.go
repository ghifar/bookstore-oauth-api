package app

import (
	"github.com/ghifar/bookstore-oauth-api/clients/cassandra"
	"github.com/ghifar/bookstore-oauth-api/domain/access_token"
	"github.com/ghifar/bookstore-oauth-api/domain/db_repository/db"
	"github.com/ghifar/bookstore-oauth-api/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()

	tokenService := access_token.NewService(db.NewRepository())
	tokenHandler := http.NewHandler(tokenService)

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.Run(":8080")

}
