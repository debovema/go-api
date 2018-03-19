package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"time"
)

type AuthType int

const (
	Basic AuthType = iota
	JWT
)

func AuthorizationHandler(authType AuthType, group *gin.RouterGroup, accounts gin.Accounts) gin.HandlerFunc {
	switch authType {
	case Basic:
		return gin.BasicAuth(accounts)

	case JWT:
		{
			// the jwt middleware
			authMiddleware := &jwt.GinJWTMiddleware{
				Realm:      "test zone",
				Key:        []byte("secret key"),
				Timeout:    time.Minute,
				MaxRefresh: time.Hour,
				Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
					for k, v := range accounts {
						if userId == k && password == v {
							return userId, true
						}
					}

					return userId, false
				},
				Authorizator: func(userId string, c *gin.Context) bool {
					return true // any authenticated user is authorized
				},
				Unauthorized: func(c *gin.Context, code int, message string) {
					c.JSON(code, gin.H{
						"code":    code,
						"message": message,
					})
				},
				TokenLookup:   "header:Authorization",
				TokenHeadName: "Bearer",
				TimeFunc:      time.Now,
			}

			if group != nil {
				group.POST("/login", authMiddleware.LoginHandler)
			}
			return authMiddleware.MiddlewareFunc()
		}

	default:
		return func(context *gin.Context) {
			// no auth
		}
	}
}
