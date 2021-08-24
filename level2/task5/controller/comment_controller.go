package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"task5/models"
)

type CommentController struct {
	controller *Controller
}

// CreateComment godoc
// @Summary Create comment
// @Description create a comment by sending valid JSON body
// @ID create-comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Router /create/comment [post]
func (commentController *CommentController) CreateComment(c echo.Context) error {
	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	commentController.controller.DB.GormDB.Create(comment)

	return c.JSON(http.StatusCreated, comment)
}

// GetComment godoc
// @Summary Get comment
// @Description get a comment by id
// @ID get-comment
// @Tags comments
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /get/comment/{id} [get]
func (commentController *CommentController) GetComment(c echo.Context) error {
	key := c.Param("id")
	comment := new(models.Comment)
	commentController.controller.DB.GormDB.First(comment, key)

	return c.JSON(http.StatusOK, comment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description update comment by sending valid JSON body
// @ID update-comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Router /update/comment/{id} [put]
func (commentController *CommentController) UpdateComment(c echo.Context) error {
	key := c.Param("id")
	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	comment.ID, _ = strconv.Atoi(key)
	commentController.controller.DB.GormDB.Save(comment)

	return c.JSON(http.StatusCreated, comment)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description delete comment by id
// @ID delete-comment
// @Tags comments
// @Produce  plain
// @Param id path int true "Comment ID"
// @Success 200 {plain} string
// @Router /delete/comment/{id} [delete]
func (commentController *CommentController) DeleteComment(c echo.Context) error {
	key := c.Param("id")
	id, _ := strconv.ParseInt(key, 10, 64)
	comment := new(models.Comment)
	commentController.controller.DB.GormDB.Where("id = ?", id).Delete(comment)

	return c.String(http.StatusOK, "Deleted comment with id = " + key)
}
