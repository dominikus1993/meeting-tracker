package dto

import "time"

type StartedMeetingDto struct {
	MeetingID string    `json:"meetingId,omitempty"`
	Leader    string    `json:"leader,omitempty"`
	Start     time.Time `json:"start,omitempty"`
	End       time.Time `json:"end,omitempty"`
	Finished  bool
}
