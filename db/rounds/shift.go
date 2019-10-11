package rounds

import (
	"github.com/uht-hack/unsure/db/events"
	"github.com/luno/shift"
)

//go:generate shiftgen -inserter=joinReq -updaters=joinedReq,collectReq,collectedReq,submitReq,submittedReq,successReq,failedReq -table=engine_rounds

var fsm = shift.NewFSM(events.GetTable()).
	Insert(RoundStatusJoin, joinReq{},
		RoundStatusJoined, RoundStatusFailed).
	Update(RoundStatusJoined, joinedReq{},
		RoundStatusJoined, RoundStatusCollect, RoundStatusFailed).
	Update(RoundStatusCollect, collectReq{},
		RoundStatusCollected, RoundStatusFailed).
	Update(RoundStatusCollected, collectedReq{},
		RoundStatusCollected, RoundStatusSubmit, RoundStatusFailed).
	Update(RoundStatusSubmit, submitReq{},
		RoundStatusSubmitted, RoundStatusFailed).
	Update(RoundStatusSubmitted, submittedReq{},
		RoundStatusSubmitted, RoundStatusSuccess, RoundStatusFailed).
	Update(RoundStatusSuccess, successReq{}).
	Update(RoundStatusFailed, failedReq{}).
	Build()

type joinReq struct {
	MatchID int64
	Index   int
	Team    string
}

type joinedReq struct {
	ID    int64
	State RoundStatus
}

type collectReq struct {
	ID    int64
	State RoundState
}

type collectedReq struct {
	ID    int64
	State RoundState
}

type submitReq struct {
	ID int64
}

type submittedReq struct {
	ID    int64
	State RoundState
}

type successReq struct {
	ID int64
}

type failedReq struct {
	ID    int64
	Error string
}