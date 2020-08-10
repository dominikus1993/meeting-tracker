package model

import (
	"time"
)

type Meeting struct {
	MeetingID string
	Leaders   []string
	Start     time.Time
	Finish    time.Time
	Finished  bool
}

func NewMeeting(id string, leaders []string) *Meeting {
	now := time.Now()
	return &Meeting{
		MeetingID: id,
		Leaders:   leaders,
		Start:     now,
		Finish:    time.Time{},
		Finished:  false,
	}
}
