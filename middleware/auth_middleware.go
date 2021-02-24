package middleware

import (
	"strings"
	"test/custom_errors"
	"test/interface/service"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware(userService service.UserService) gin.HandlerFunc {
	if userService == nil {
		panic("userService argument must not be nil")
	}

	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			respondWithError(c, 401, "No Bearer token found")
			return
		}

		token := splitToken[1]
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		claim, err := userService.ValidateToken(token)
		if err != nil {
			if _, ok := err.(*custom_errors.UnauthorizedError); ok {
				respondWithError(c, 401, err.Error())
			} else {
				respondWithError(c, 500, err.Error())
			}

			return
		}

		c.Set("Claim", claim)
		c.Next()
	}
}
