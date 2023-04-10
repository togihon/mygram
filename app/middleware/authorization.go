package middleware

import (
	"mygram/app/entity"
	"mygram/pkg/database"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(endpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := database.Connect()
		param, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, entity.ResponseFailed{
				Success: false,
				Message: "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		switch endpoint {
		case "photo":
			Entity := entity.MyGramPhoto{}
			err := db.Select("my_gram_user_id").First(&Entity, uint(param)).Error

			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, entity.ResponseFailed{
					Success: false,
					Message: "Invalid parameter",
				})
				return
			}

			if Entity.MyGramUserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ResponseFailed{
					Success: false,
					Message: "You are not allowed to access this data",
				})
			}
		case "comment":
			Entity := entity.MyGramComment{}
			err := db.Select("my_gram_user_id").First(&Entity, uint(param)).Error

			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, entity.ResponseFailed{
					Success: false,
					Message: "Data not found or exist",
				})
				return
			}

			if Entity.MyGramUserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ResponseFailed{
					Success: false,
					Message: "You are not allowed to access this data",
				})
			}
		case "socialMedia":
			Entity := entity.MyGramSocialMedia{}
			err := db.Select("my_gram_user_id").First(&Entity, uint(param)).Error

			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, entity.ResponseFailed{
					Success: false,
					Message: "Data not found or exist",
				})
				return
			}

			if Entity.MyGramUserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ResponseFailed{
					Success: false,
					Message: "You are not allowed to access this data",
				})
			}
		default:

		}

	}
}
