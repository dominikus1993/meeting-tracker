package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoMeeting struct {
	MeetingID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}