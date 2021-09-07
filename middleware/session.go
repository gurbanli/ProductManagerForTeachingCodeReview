package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gurbanli/ProductManager/config"
	"net/http"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")
		if err != nil {
			c.Abort()
		}
		if session.Values["authenticated"] != true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}
	}
}
