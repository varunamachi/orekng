package rest

import (
	"time"
)

//Result - holds result message of a non GET api
type Result struct {
	Operation string `json:"operation"`
	Message   string `json:"message"`
	Error     string `json:"error"`
}

//Session - represents a session object
type Session struct {
	SessionID string            `json:"sessionId" db:"session_id"`
	UserID    string            `json:"userId" db:"user"`
	StartTime time.Time         `json:"startTime" db:"start_time"`
	Props     map[string]string `json:"properties"`
}
