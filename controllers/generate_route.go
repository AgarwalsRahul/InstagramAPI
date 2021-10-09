// GenerateUserRoutes :- Helper function for collating user routes
package controllers

import (
	"github.com/AgarwalsRahul/InstagramAPI/router"
)

type Controller struct {
	UserController *UserController
	PostController *PostController
}

func NewController(uc *UserController, pc *PostController) *Controller {
	return &Controller{uc, pc}
}

func (c Controller) GenerateRoutes() []router.Route {
	// Create user setup
	createUserRoute := router.Route{
		Name:            "Create User",
		Method:          "POST",
		Path:            "/users",
		HandlerFunction: c.UserController.CreateUser,
	}

	// Get user setup
	getUserRoute := router.Route{
		Name:            "Get User",
		Method:          "GET",
		Path:            "/users/:id",
		HandlerFunction: c.UserController.GetUser,
	}

	getPostByIdRoute := router.Route{
		Name:            "Get PosT",
		Method:          "GET",
		Path:            "/post/:id",
		HandlerFunction: c.PostController.GetPostById,
	}
	createPost := router.Route{
		Name:            "Create Post",
		Method:          "POST",
		Path:            "/posts",
		HandlerFunction: c.PostController.CreatePost,
	}

	getAllPosts := router.Route{
		Name:            "Get All Posts",
		Method:          "GET",
		Path:            "/posts/users/:id/:page",
		HandlerFunction: c.PostController.GetAllPosts,
	}

	// collate all routes
	routes := []router.Route{createUserRoute, getUserRoute, createPost, getAllPosts, getPostByIdRoute}

	return routes
}
