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

// MyGramGetAllSocialMedia godoc
// @Summary Get all social media
// @Description User can retrieve all social media and no need to login
// @Tags social-medias
// @Consumes ({mpfd,json})
// @Produce json
// @Success 200 {object} entity.Response "Will send all social media datas"
// @Failure 404  {object}  entity.Response "If there is no social media, error will appear"
// @Router /social-media [GET]
func MyGramGetAllSocialMedia(c *gin.Context) {
	db, _ := database.Connect()
	SocialMedia := []entity.MyGramSocialMedia{}
	err := db.Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusNotFound, entity.Response{
			Success: false,
			Message: "There's no SocialMedia found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Social medias has been loaded successfully",
		Data:    SocialMedia,
	})
}

// MyGramGetSocialMedia godoc
// @Summary Get one social media
// @Description User can retrieve a social media and doesn't need to login
// @Tags social-medias
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "social media id"
// @Success 200 {object} entity.Response "If a social media's id matches with the parameter"
// @Failure 404  {object}  entity.Response "If the social media's id doesn't match with the parameter, error will appear"
// @Router /social-media/{id} [GET]
func MyGramGetSocialMedia(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	SocialMedia := entity.MyGramSocialMedia{}

	//get parameter
	socialMediaID, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	//query select * from SocialMedia where id = param
	err := db.First(&SocialMedia, "id = ?", socialMediaID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, entity.Response{
			Success: false,
			Message: "Social media not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Social media has been loaded successfully",
		Data:    SocialMedia,
	})
}

// MyGramCreateSocialMedia godoc
// @Summary Create a social media
// @Description User can create a social media.
// @Tags social-medias
// @Consumes ({mpfd,json})
// @Produce json
// @Param name formData string true "social media name"
// @Param social_media_url formData string true "social media url"
// @Success 201 {object} entity.Response "If all of the parameters filled and you are logged in"
// @Failure 401  {object}  entity.Response "If you are not login or some parameters not filled, error will appear"
// @Security Bearer
// @Router /social-media [POST]
func MyGramCreateSocialMedia(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := entity.MyGramSocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.MyGramUserID = userID
	err := db.Debug().Create(&SocialMedia).Error

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
		Message: "Social media has been created successfully",
		Data:    SocialMedia,
	})
}

// MyGramUpdateSocialMedia godoc
// @Summary Edit a social media
// @Description User can edit their own social media.
// @Tags social-medias
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "social media id"
// @Param name formData string true "social media name"
// @Param social_media_url formData string true "social media url"
// @Success 200 {object} entity.Response "If all the parameters are valid"
// @Failure 400  {object}  entity.Response "If there is something wrong, error will appear"
// @Security Bearer
// @Router /social-media/{id} [PUT]
func MyGramUpdateSocialMedia(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := entity.MyGramSocialMedia{}

	socialMediaID, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.MyGramUserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(entity.MyGramSocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Social media has been updated successfully",
		Data:    SocialMedia,
	})
}

// MyGramDeleteSocialMedia godoc
// @Summary Delete a social media
// @Description User can delete their own social media.
// @Tags social-medias
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "social media id"
// @Success 200 {object} entity.Response "If social media is exist and it's your own social media"
// @Failure 400  {object}  entity.Response "If social media's id is not your own or if the comment doesn't exist, error will appear"
// @Security Bearer
// @Router /social-media/{id} [DELETE]
func MyGramDeleteSocialMedia(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	SocialMedia := entity.MyGramSocialMedia{}

	//get parameter
	socialMediaID, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Where("id = ?", socialMediaID).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Sosial media berhasil dihapus",
		Data:    nil,
	})
}
