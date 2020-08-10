package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/domain/model"
	"time"
)

type MongoMeeting struct {
	MeetingID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Leaders   []string           `json:"leaders,omitempty" bson:"leader,omitempty"`
	Start     time.Time          `json:"start,omitempty" bson:"start,omitempty"`
	Finish    time.Time          `json:"finish,omitempty" bson:"finish,omitempty"`
	Finished  bool               `json:"finished,omitempty" bson:"finished,omitempty"`
}

func (m *MongoMeeting) ToDomainMeeting() *model.Meeting {
	return &model.Meeting{
		MeetingID: m.MeetingID.String(),
		Leaders:   m.Leaders,
		Start:     m.Start,
		Finish:    m.Finish,
		Finished:  m.Finished,
	}
}
