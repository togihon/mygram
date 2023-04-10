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
// @Success 201 {object} entity.ResponseRegister "Jika semua field benar, maka akun akan dibuat "
// @Failure 400  {object}  entity.ResponseFailed "Jika terdapat kesalahan akan muncul error"
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
		c.JSON(http.StatusBadRequest, entity.ResponseFailed{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, entity.ResponseRegister{
		Success: true,
		Message: "Account has been created successfully",
		Data: struct {
			ID    uint   "json:\"id\" example:\"1\""
			Email string "json:\"email\" example:\"user@mail.com\""
			Uname string "json:\"username\" example:\"user\""
			Age   int    "json:\"age\" example:\"18\""
		}{ID: User.ID, Email: User.Email, Uname: User.Username, Age: int(User.Age)},
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
// @Success 200 {object} entity.ResponseLogin "Jika email dan password benar, maka akan mendapatkan token"
// @Failure 401  {object}  entity.ResponseFailed "Jika email dan password salah, maka akan muncul error"
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
			entity.ResponseFailed{
				Success: false,
				Message: "Invalid email or password",
			})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, entity.ResponseFailed{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, entity.ResponseLogin{
		Success: true,
		Message: "User logged in successfully",
		Data: struct {
			Token string "json:\"token\" example:\"eyJhbGciOiJI....\""
		}{Token: token},
	})
}
