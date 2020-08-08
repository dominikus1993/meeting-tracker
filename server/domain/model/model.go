package model

import (
	"time"
)

type Meeting struct {
	MeetingID string
	Leader    string
	Start     time.Time
	Finish    time.Time
	Finished  bool
}

func NewMeeting(id, leader string) *Meeting {
	now := time.Now()
	return &Meeting{
		MeetingID: id,
		Leader:    leader,
		Start:     now,
		Finish:    time.Time{},
		Finished:  false,
	}
}
