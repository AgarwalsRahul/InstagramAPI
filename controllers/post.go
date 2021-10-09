package controllers

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"strconv"
	"time"

	appError "github.com/AgarwalsRahul/InstagramAPI/errors"
	"github.com/AgarwalsRahul/InstagramAPI/models"
	"github.com/AgarwalsRahul/InstagramAPI/response"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct {
	session *mgo.Session
}

const PAGINATION_SIZE int64 = 30

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

func (pc PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := models.Post{}
	// json.NewDecoder(r.Body).Decode(&post)
	err := decode(r, &post)

	if err != nil {
		validationData := appError.Params{
			"caption":  "required",
			"imageUrl": "required",
			"userId":   "required",
		}
		errResp := appError.NewAPIError(http.StatusBadRequest, "BAD_REQUEST", "Please provide valid user data.", validationData)
		appError.WriteErrorResponse(w, errResp)
		return
	}
	post.Id = bson.NewObjectId()
	post.PostedTimestamp = time.Now().Format("01-02-2006 15:04:05")
	if err := pc.session.DB("instagram").C("posts").Database.C("users").Database.C(post.UserId).Insert(post); err != nil {
		apiError := appError.InternalServerError(err)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	resp := response.GenericResponse(http.StatusCreated, "Post created successfully", post)
	response.WriteResponse(w, resp)
	// w.Header().Set("content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "%s\n", postJson)
}

func (pc PostController) GetPostById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	userId := r.Header.Get("Authorization")
	if !bson.IsObjectIdHex(id) {
		apiError := appError.NotFound("Invalid ID " + id)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	objectId := bson.ObjectIdHex(id)
	post := models.Post{}
	if err := pc.session.DB("instagram").C("posts").Database.C("user").Database.C(userId).FindId(objectId).One(&post); err != nil {
		apiError := appError.InternalServerError(err)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	// postJson, err := json.Marshal(post)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Header().Set("content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%s\n", postJson)
	resp := response.GenericResponse(http.StatusFound, "Post found successfully", post)
	response.WriteResponse(w, resp)
}

func (pc PostController) GetAllPosts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	page := params.ByName("page")
	pageInt, err := strconv.ParseInt(page, 0, 64)
	if err != nil {
		apiError := appError.NotFound("Invalid Page " + page)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	skipPages := (pageInt - 1) * PAGINATION_SIZE
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	if !bson.IsObjectIdHex(id) {
		// w.WriteHeader(http.StatusNotFound)
		apiError := appError.NotFound("Invalid ID " + id)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	objectId := bson.ObjectIdHex(id)
	var posts []models.Post
	if err := pc.session.DB("instagram").C("posts").Database.C("users").Database.C(objectId.Hex()).Find(bson.D{}).Skip(int(skipPages)).Limit(int(pageInt * PAGINATION_SIZE)).All(&posts); err != nil {
		apiError := appError.InternalServerError(err)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	// postJson, err := json.Marshal(posts)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Header().Set("content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%s\n", postJson)
	resp := response.GenericResponse(http.StatusFound, "Post retrieved successfully", posts)
	response.WriteResponse(w, resp)
}
