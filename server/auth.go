package server

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	authmiddleware "github.com/udaya2899/go-gin-starter/middleware/authMiddleware"
)

func setAuthRoutes(r *gin.Engine) {

	authm, err := authmiddleware.Get()
	if err != nil {
		logrus.Panicf("Cannot initiate auth middleware: err: %v", err)

	}

	r.POST("/login", authm.LoginHandler)

	r.NoRoute(authm.MiddlewareFunc())

	auth := r.Group("/auth")

	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authm.RefreshHandler)
	auth.Use(authm.MiddlewareFunc())
	{
		auth.GET("/hello", getAuthStatus)
	}
}

func getAuthStatus(c *gin.Context) {
	c.String(http.StatusOK, "authed_pong")
}

func noRoute(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
