package controllers

import (
	"encoding/json"
	"errors"
	// "fmt"
	"net/http"

	appError "github.com/AgarwalsRahul/InstagramAPI/errors"
	"github.com/AgarwalsRahul/InstagramAPI/models"
	"github.com/AgarwalsRahul/InstagramAPI/response"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ok represents types capable of validating
// themselves.
type ok interface {
	OK() error
}
type UserController struct {
	session *mgo.Session
}

func NewUserConroller(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if !bson.IsObjectIdHex(id) {
		// w.WriteHeader(http.StatusNotFound)
		apiError := appError.NotFound("Invalid ID " + id)
		appError.WriteErrorResponse(w, apiError)
		return
	}

	objectId := bson.ObjectIdHex(id)
	user := models.User{}
	if err := uc.session.DB("instagram").C("users").FindId(objectId).One(&user); err != nil {
		apiError := appError.NotFound("Invalid ID " + id)
		appError.WriteErrorResponse(w, apiError)
		return
	}
	resp := response.GenericResponse(http.StatusFound, "User found successfully", user)
	response.WriteResponse(w, resp)

	// userJson, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Header().Set("content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := decode(r, &user)
	// json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		validationData := appError.Params{
			"Name":     "required",
			"password": "required",
			"email":    "required",
		}
		errResp := appError.NewAPIError(http.StatusBadRequest, "BAD_REQUEST", "Please provide valid user data.", validationData)
		appError.WriteErrorResponse(w, errResp)
		return
	}

	user.Id = bson.NewObjectId()
	uc.session.DB("instagram").C("users").Insert(user)
	resp := response.GenericResponse(http.StatusCreated, "User Created Successfully.", user)
	response.WriteResponse(w, resp)
	// userJson, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Header().Set("content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "%s\n", userJson)
}

// decode can be this simple to start with, but can be extended
// later to support different formats and behaviours without
// changing the interface.
func decode(r *http.Request, v ok) error {
	if r.Body == nil {
		return errors.New("Invalid Body")
	}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}
