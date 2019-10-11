package matches

import (
	"github.com/luno/shift"
)

//go:generate shiftgen -inserter=startReq -updaters=endReq -table=matches

//                 | --> Success
// New --> Started |
//                 L --> Failed

var fsm = shift.NewFSM(events).
	Insert(MatchStatusStarted, startReq{}, MatchStatusEnded).
	Update(MatchStatusEnded, endReq{}).
	Build()

type startReq struct {
	Players int
}

type endReq struct {
	ID      int64
}
