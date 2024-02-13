// platform/router/router.go

package router

import (
	"encoding/gob"
	"github.com/jairogloz/go-auth0/platform/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/jairogloz/go-auth0/platform/authenticator"
	"github.com/jairogloz/go-auth0/platform/web/app/callback"
	"github.com/jairogloz/go-auth0/platform/web/app/login"
	"github.com/jairogloz/go-auth0/platform/web/app/logout"
	"github.com/jairogloz/go-auth0/platform/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "platform/web/static")
	router.LoadHTMLGlob("platform/web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
