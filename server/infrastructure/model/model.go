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
		MeetingID: m.MeetingID.Hex(),
		Leaders:   m.Leaders,
		Start:     m.Start,
		Finish:    m.Finish,
		Finished:  m.Finished,
	}
}

func FromDomainMeeting(meeting *model.Meeting) (*MongoMeeting, error) {
	id, err := primitive.ObjectIDFromHex(meeting.MeetingID)
	if err != nil {
		return nil, err
	}
	return &MongoMeeting{
		MeetingID: id,
		Leaders:   meeting.Leaders,
		Start:     meeting.Start,
		Finish:    meeting.Finish,
		Finished:  meeting.Finished,
	}, nil
}

func FromDomainMeetingAsFinished(meeting *model.Meeting) (*MongoMeeting, error) {
	id, err := primitive.ObjectIDFromHex(meeting.MeetingID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &MongoMeeting{
		MeetingID: id,
		Leaders:   meeting.Leaders,
		Start:     meeting.Start,
		Finish:    now,
		Finished:  true,
	}, nil
}
