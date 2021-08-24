package controller

import "task5/store"

type Controller struct {
	DB *store.Database
	postController *PostController
	commentController *CommentController
}

func (controller *Controller) PostController() *PostController {
	if controller.postController != nil {
		return controller.postController
	}

	controller.postController = &PostController{controller: controller}

	return controller.postController
}

func (controller *Controller) CommentController() *CommentController {
	if controller.commentController != nil {
		return controller.commentController
	}

	controller.commentController = &CommentController{controller: controller}

	return controller.commentController
}

func New() *Controller {
	return &Controller{DB: store.GetDB()}
}
