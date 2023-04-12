package services

import (
	"mygram/app/entity"
	"mygram/pkg/database"
	"mygram/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MyGramGetAllPhoto godoc
// @Summary Get all photos
// @Description User can retrieve all photos and no need to login
// @Tags photos
// @Consumes ({mpfd,json})
// @Produce json
// @Success 200 {object} entity.Response "Will send all photo datas"
// @Failure 404  {object}  entity.ResponseFailed "If there is no photo, error will appear"
// @Router /photos [GET]
func MyGramGetAllPhoto(c *gin.Context) {
	db, _ := database.Connect()
	Photo := []entity.MyGramPhoto{}
	err := db.Order("created_at desc").Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, entity.ResponseFailed{
			Success: false,
			Message: "There's no photo found",
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Photos has been loaded successfully",
		Data:    Photo,
	})
}

// MyGramGetPhoto godoc
// @Summary Get one photo
// @Description User can retrieve a photo and no need to login
// @Tags photos
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "photo id"
// @Success 200 {object} entity.Response "If a photo's id matches with the parameter"
// @Failure 404  {object}  entity.ResponseFailed "If the photo's id doesn't match with the parameter, error will appear"
// @Router /photos/{id} [GET]
func MyGramGetPhoto(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	Photo := entity.MyGramPhoto{}

	//get parameter
	photoID, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	//query select * from photo where id = param
	err := db.First(&Photo, "id = ?", photoID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, entity.ResponseFailed{
			Success: false,
			Message: "Photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Photo has been loaded successfully",
		Data:    Photo,
	})
}

// MyGramCreatePhoto godoc
// @Summary Upload a photo
// @Description User can upload a photo.
// @Tags photos
// @Consumes ({mpfd,json})
// @Produce json
// @Param title formData string true "photo title"
// @Param caption formData string true "photo caption"
// @Param photo_url formData string true "photo url"
// @Success 201 {object} entity.Response "If all of the parameters filled and you're logged in"
// @Failure 404  {object}  entity.ResponseFailed "If you are not login or some parameters not filled, error will appear"
// @Security Bearer
// @Router /photos [POST]
func MyGramCreatePhoto(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := entity.MyGramPhoto{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.MyGramUserID = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResponseFailed{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, entity.Response{
		Success: true,
		Message: "Photo has been created successfully",
		Data:    Photo,
	})
}

// MyGramUpdatePhoto godoc
// @Summary Edit a photo
// @Description User can edit their own photo.
// @Tags photos
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "photo id"
// @Param title formData string true "photo title"
// @Param caption formData string true "photo caption"
// @Param photo_url formData string true "photo url"
// @Success 200 {object} entity.Response "If the parameters are valid"
// @Failure 401  {object}  entity.ResponseFailed "If there is something wrong, error will appear"
// @Security Bearer
// @Router /photos/{id} [PUT]
func MyGramUpdatePhoto(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := entity.MyGramPhoto{}

	photoID, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.MyGramUserID = userID
	Photo.ID = uint(photoID)

	err := db.Model(&Photo).Where("id = ?", photoID).Updates(entity.MyGramPhoto{Title: Photo.Title, Caption: Photo.Caption, Photo_URL: Photo.Photo_URL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResponseFailed{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Photo has been updated successfully",
		Data:    Photo,
	})

}

// MyGramDeletePhoto godoc
// @Summary Delete a photo
// @Description User can delete their own photo.
// @Tags photos
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "photo id"
// @Success 200 {object} entity.Response "If photo is exist and it's your own photo"
// @Failure 400  {object}  entity.ResponseFailed "If the photo is not your own or if the photo doesn't exist, error will appear"
// @Security Bearer
// @Router /photos/{id} [DELETE]
func MyGramDeletePhoto(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	Photo := entity.MyGramPhoto{}

	//get parameter
	photoID, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("id = ?", photoID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResponseFailed{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.ResponseSuccess{
		Success: true,
		Message: "Photo has been deleted successfully",
	})
}
