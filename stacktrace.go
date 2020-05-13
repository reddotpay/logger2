package logger2

import "github.com/pkg/errors"

// StackTrace structure for stacktracing
type StackTrace struct {
	Type    string      `json:"type"`
	Summary string      `json:"summary,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
	Caller  interface{} `json:"caller"`
}

// StackTraceArray list of StackTrace
type StackTraceArray []*StackTrace

type stackTracer interface {
	StackTrace() errors.StackTrace
}
