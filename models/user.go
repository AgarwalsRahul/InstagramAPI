package models

import "gopkg.in/mgo.v2/bson"

type missingFieldError string

func (m missingFieldError) Error() string {
	return string(m) + " is required"
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

// UserStore :- This is a dictionary of users
var UserStore = make(map[string]User)

// OK represents types capable of validating themselves.
// This would be utilized when decoding from request body.
// Here, we can decide to check for required fields as needed.
func (u *User) OK() error {
	if len(u.Name) == 0 {
		return missingFieldError("Name")
	}
	if len(u.Email) == 0 {
		return missingFieldError("email")
	}

	if len(u.Password) == 0 {
		return missingFieldError("password")
	}
	return nil
}
