package rounds

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"sort"
	"time"

	"github.com/corverroos/unsure/engine"
)

type Round struct {
	ID        int64
	MatchID   int64
	Index     int64
	Team      string
	Status    RoundStatus
	State     RoundState
	Error     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//go:generate stringer -type=RoundStatus -trimprefix=RoundStatus

type RoundStatus int

func (rs RoundStatus) Enum() int {
	return int(rs)
}

func (rs RoundStatus) Valid() bool {
	return rs > RoundStatusUnknown && rs < roundStatusSentinel
}

func (rs RoundStatus) ShiftStatus() {}

func (rs RoundStatus) ReflexType() int {
	return engine.RoundEventOffset + int(rs) // Hack to combine Match and Round events in same table.
}

const (
	RoundStatusUnknown   RoundStatus = 0
	RoundStatusJoin      RoundStatus = 1
	RoundStatusJoined    RoundStatus = 2
	RoundStatusCollect   RoundStatus = 3
	RoundStatusCollected RoundStatus = 4
	RoundStatusSubmit    RoundStatus = 5
	RoundStatusSubmitted RoundStatus = 6
	RoundStatusSuccess   RoundStatus = 7
	RoundStatusFailed    RoundStatus = 8
	roundStatusSentinel  RoundStatus = 9
)

type RoundState struct {
	Players []RoundPlayerState
}

type RoundPlayerState struct {
	Name      string
	Rank      int
	Parts     map[string]int
	Included  bool
	Collected bool
	Submitted bool
}

func (rs RoundState) Value() (driver.Value, error) {
	return json.MarshalIndent(rs, "", " ")
}

func (rs *RoundState) Scan(src interface{}) error {
	var s sql.NullString
	if err := s.Scan(src); err != nil {
		return err
	}
	*rs = RoundState{}
	if !s.Valid {
		return nil
	}
	return json.Unmarshal([]byte(s.String), rs)
}

func (rs RoundState) GetPlayer(player string) (int, RoundPlayerState, bool) {
	for i, m := range rs.Players {
		if m.Name == player {
			return i, m, true
		}
	}
	return 0, RoundPlayerState{}, false
}

func (rs RoundState) GetSubmitOrder() []RoundPlayerState {
	var res []RoundPlayerState
	for _, m := range rs.Players {
		if !m.Included {
			continue
		}
		res = append(res, m)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Rank < res[j].Rank
	})

	return res
}

func (rs RoundState) GetTotal(player string) int {
	var res int
	for _, m := range rs.Players {
		res += m.Parts[player]
	}
	return res
}