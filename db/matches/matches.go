package matches

type Match struct {
	id int64
	status int
	players int
}

type MatchStatus int

func (m MatchStatus) Enum() int {
	return int(m)
}

func (m MatchStatus) ShiftStatus() {}

func (m MatchStatus) ReflexType() int {
	return int(m)
}

const (
	MatchStatusUnknown  MatchStatus = 0
	MatchStatusStarted  MatchStatus = 1
	MatchStatusEnded    MatchStatus = 2
	matchStatusSentinel MatchStatus = 3 // This may not
)
