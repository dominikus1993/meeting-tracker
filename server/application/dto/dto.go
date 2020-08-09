package dto

import "time"

type StartedMeetingDto struct {
	MeetingID string    `json:"meetingId,omitempty"`
	Leader    string    `json:"leader,omitempty"`
	Start     time.Time `json:"start,omitempty"`
	Finished  bool      `json:"finished,omitempty"`
}

type FinishedMeetingDto struct {
	MeetingID string    `json:"meetingId,omitempty"`
	Leader    string    `json:"leader,omitempty"`
	Start     time.Time `json:"start,omitempty"`
	End       time.Time `json:"finish,omitempty"`
	Finished  bool      `json:"finished,omitempty"`
}

type MeetingDto struct {
	MeetingID string    `json:"meetingId,omitempty"`
	Leader    string    `json:"leader,omitempty"`
	Start     time.Time `json:"start,omitempty"`
	End       time.Time `json:"end,omitempty"`
	Finished  bool      `json:"finished,omitempty"`
}
