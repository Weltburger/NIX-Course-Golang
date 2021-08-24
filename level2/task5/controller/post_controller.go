package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"task5/models"
)

type PostController struct {
	controller *Controller
}

// CreatePost godoc
// @Summary Create post
// @Description create a post by sending valid JSON body
// @ID create-post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param comment body models.Post true "Post"
// @Success 200 {object} models.Post
// @Router /create/post [post]
func (postController *PostController) CreatePost(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	postController.controller.DB.GormDB.Create(post)

	return c.JSON(http.StatusCreated, post)
}

// GetPost godoc
// @Summary Get post
// @Description get a post by id
// @ID get-post
// @Tags posts
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /get/post/{id} [get]
func (postController *PostController) GetPost(c echo.Context) error {
	key := c.Param("id")
	post := &models.Post{}
	postController.controller.DB.GormDB.First(post, key)

	return c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update post
// @Description update post by sending valid JSON body
// @ID update-post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param comment body models.Post true "Post"
// @Success 200 {object} models.Post
// @Router /update/post/{id} [put]
func (postController *PostController) UpdatePost(c echo.Context) error {
	key := c.Param("id")
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	post.ID, _ = strconv.Atoi(key)
	postController.controller.DB.GormDB.Save(post)

	return c.JSON(http.StatusCreated, post)
}

// DeletePost godoc
// @Summary Delete post
// @Description delete post by id
// @ID delete-post
// @Tags posts
// @Produce  plain
// @Param id path int true "Post ID"
// @Success 200 {plain} string
// @Router /delete/post/{id} [delete]
func (postController *PostController) DeletePost(c echo.Context) error {
	key := c.Param("id")
	id, _ := strconv.ParseInt(key, 10, 64)
	post := new(models.Post)
	postController.controller.DB.GormDB.Where("id = ?", id).Delete(post)

	return c.String(http.StatusOK, "Deleted post with id = " + key)
}
