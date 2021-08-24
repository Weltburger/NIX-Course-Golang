package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"task5/controller"
	_ "task5/docs"
)

// @title NIX task5 API
// @version 1.0
// @description This is a sample server post and comment server.
// @termsOfService http://swagger.io/terms/

// @contact.name Weltburger
// @contact.url https://github.com/Weltburger
// @contact.email l.weltburger.l@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /


func main() {
	control:= controller.New()

	e := echo.New()

	e.POST("/create/post", control.PostController().CreatePost)
	e.GET("/get/post/:id", control.PostController().GetPost)
	e.PUT("/update/post/:id", control.PostController().UpdatePost)
	e.DELETE("/delete/post/:id", control.PostController().DeletePost)

	e.POST("/create/comment", control.CommentController().CreateComment)
	e.GET("/get/comment/:id", control.CommentController().GetComment)
	e.PUT("/update/comment/:id", control.CommentController().UpdateComment)
	e.DELETE("/delete/comment/:id", control.CommentController().DeleteComment)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
