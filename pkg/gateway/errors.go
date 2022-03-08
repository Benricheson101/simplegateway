package gateway

import (
	"errors"
)

var (
	ErrMissingGatewayURL     = errors.New("gateway url missing")
	ErrAlreadyOpenConnection = errors.New("there is already an open connection")
	ErrNilConnection         = errors.New("no active websocket connection")
	ErrCouldNotConnect       = errors.New("failed to connect to discord")
	ErrNilHeartbeatInterval  = errors.New("heartbeat interval not set")
	ErrNoDispatchEventStruct = errors.New("no corresponding struct for dispatched event")
	ErrMissingResumeFields   = errors.New("missing one or more fields required for resume")
	ErrResumeFail            = errors.New("failed to resume gateway connection")
	ErrInvalidSession        = errors.New("invalid session")
	ErrNotReady              = errors.New("unexpected non READY msg")
)
