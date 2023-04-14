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

// MyGramGetAllComment godoc
// @Summary Get all comments
// @Description User can retrieve all comments and no need to login
// @Tags comments
// @Consumes ({mpfd,json})
// @Produce json
// @Success 200 {object} entity.Response "Will send all comments"
// @Failure 404  {object}  entity.Response "If there is no comment, error will appear"
// @Router /comments [GET]
func MyGramGetAllComment(c *gin.Context) {
	db, _ := database.Connect()
	Comment := []entity.MyGramComment{}
	err := db.Order("created_at desc").Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusNotFound, entity.Response{
			Success: false,
			Message: "There's no comment found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Comments has been loaded successfully",
		Data:    Comment,
	})
}

// MyGramGetComment godoc
// @Summary Get one comment
// @Description User can retrieve a comment and no need to login
// @Tags comments
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "comment id"
// @Success 200 {object} entity.Response "If a comment's id matches with the parameter"
// @Failure 404  {object}  entity.Response "If the comments's id doesn't match with the parameter, error will appear"
// @Router /comments/{id} [GET]
func MyGramGetComment(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	Comment := entity.MyGramComment{}

	//get parameter
	commentID, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	//query select * from comment where id = param
	err := db.First(&Comment, "id = ?", commentID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, entity.Response{
			Success: false,
			Message: "Comment not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Comment has been loaded successfully",
		Data:    Comment,
	})
}

// MyGramCreateComment godoc
// @Summary Create a comment
// @Description User can create a comment.
// @Tags comments
// @Consumes ({mpfd,json})
// @Produce json
// @Param photo_id formData int true "photo id"
// @Param message formData string true "your comment"
// @Success 201 {object} entity.Response "If all of the parameters filled and you're login"
// @Failure 404 {object} entity.Response "If photo id's not found"
// @Failure 401  {object}  entity.Response "If you are not login or some parameters not filled, error will appear"
// @Security Bearer
// @Router /comments [POST]
func MyGramCreateComment(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := entity.MyGramComment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.MyGramUserID = userID
	err := db.Debug().Create(&Comment).Error

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
		Message: "Comment has been created successfully",
		Data:    Comment,
	})
}

// MyGramUpdateComment godoc
// @Summary Edit a comment
// @Description User can edit their own comment.
// @Tags comments
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "comment id"
// @Param message formData string true "your comment"
// @Success 200 {object} entity.Response "If all the parameters are valid"
// @Failure 404  {object}  entity.Response "If there is something wrong, error will appear"
// @Security Bearer
// @Router /comments/{id} [PUT]
func MyGramUpdateComment(c *gin.Context) {
	db, _ := database.Connect()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := entity.MyGramComment{}

	commentID, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.MyGramUserID = userID
	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).Updates(entity.MyGramComment{Message: Comment.Message}).Error

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
		Message: "Comment has been updated successfully",
		Data:    Comment,
	})
}

// MyGramDeleteComment godoc
// @Summary Delete a comment
// @Description User can delete their own comment.
// @Tags comments
// @Consumes ({mpfd,json})
// @Produce json
// @Param id path int true "comment id"
// @Success 200 {object} entity.Response "If comment is exist and it's your own comment"
// @Failure 400  {object}  entity.Response "If the comment's id is not your own and if the comment doesn't exist, error will appear"
// @Security Bearer
// @Router /comments/{id} [DELETE]
func MyGramDeleteComment(c *gin.Context) {
	db, _ := database.Connect()
	contentType := helpers.GetContentType(c)
	Comment := entity.MyGramComment{}

	//get parameter
	commentID, _ := strconv.Atoi(c.Param("id"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("id = ?", commentID).Delete(&Comment).Error

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
		Message: "Comment has been deleted successfully",
		Data:    nil,
	})
}
