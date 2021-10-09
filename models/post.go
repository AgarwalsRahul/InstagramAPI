package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Caption         string        `json:"caption" bson:"caption"`
	ImageUrl        string        `json:"imageUrl" bson:"imageUrl"`
	PostedTimestamp string        `json:"posted_timestamp" bson:"posted_timestamp"`
	UserId          string        `json:"userId" bson:"userId"`
}

func (p *Post) OK() error {
	if len(p.Caption) == 0 {
		return missingFieldError("Caption")
	}
	if len(p.ImageUrl) == 0 {
		return missingFieldError("ImageUrl")
	}

	if len(p.UserId) == 0 {
		return missingFieldError("UserId")
	}
	return nil
}
