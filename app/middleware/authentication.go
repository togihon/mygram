package middleware

import (
	"mygram/app/entity"
	"mygram/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, entity.ResponseFailed{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
