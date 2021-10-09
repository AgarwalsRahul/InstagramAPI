package main

import (
	"net/http"

	"github.com/AgarwalsRahul/InstagramAPI/controllers"
	"github.com/AgarwalsRahul/InstagramAPI/router"
	// "github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// r := httprouter.New()
	uc := controllers.NewUserConroller(getSession())
	pc := controllers.NewPostController(getSession())
	c := controllers.NewController(uc, pc)
	routes := c.GenerateRoutes()
	router := router.NewRouter(routes)
	// r.GET("/users/:id", uc.GetUser)
	// r.POST("/users", uc.CreateUser)
	// r.POST("/posts", pc.CreatePost)
	// r.GET("/post/:id", pc.GetPostById)
	// r.GET("/posts/users/:id/:page", pc.GetAllPosts)
	http.ListenAndServe("localhost:8080", router)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}
	return s

}
