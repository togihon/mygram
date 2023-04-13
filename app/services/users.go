package services

import (
	"mygram/app/entity"
	"mygram/pkg/database"
	"mygram/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// MyGramUserRegister godoc
// @Summary User Register
// @Description Register an account
// @Tags users
// @Consumes ({mpfd,json})
// @Produce json
// @Param email formData string true "User's email"
// @Param username formData string true "User's username"
// @Param password formData string true "User's password"
// @Param age formData int true "User's age"
// @Success 201 {object} entity.Response "If all field filled and correct, account will created "
// @Failure 400  {object}  entity.Response "If there is an error, data will set to nil"
// @Router /users/register [post]
func MyGramUserRegister(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := entity.MyGramUser{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err = db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, entity.Response{
		Success: true,
		Message: "Account has been created successfully",
		Data: entity.DataRegister{
			ID:    User.ID,
			Email: User.Email,
			Uname: User.Username,
			Age:   int(User.Age),
		},
	})

}

// MyGramUserLogin godoc
// @Summary User Login
// @Description Login to system
// @Tags users
// @Consumes ({mpfd,json})
// @Produce json
// @Param email formData string true "User's email"
// @Param password formData string true "User's password"
// @Success 200 {object} entity.Response "If email and password are correct, you will get a token"
// @Failure 401  {object}  entity.Response "If email and password are not correct, data will set to nil"
// @Router /users/login [post]
func MyGramUserLogin(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := entity.MyGramUser{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	//select data user berdasarkan email
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {

		c.JSON(http.StatusUnauthorized,
			entity.Response{
				Success: false,
				Message: "Invalid email or password",
				Data:    nil,
			})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, entity.Response{
			Success: false,
			Message: "Invalid email or password",
			Data:    nil,
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Username, User.CreatedAt)

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "User logged in successfully",
		Data: entity.DataLogin{
			Token: token,
		},
	})
}
