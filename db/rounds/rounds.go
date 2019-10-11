package rounds

import (
	"database/sql"
	"database/sql/driver"
	"context"
	"encoding/json"
	"github.com/luno/shift"
	"sort"
	"time"

	"github.com/corverroos/unsure/engine"
)

type Round struct {
	ID        int64       `protocp:"1"`
	MatchID   int64       `protocp:"2"`
	Index     int64       `protocp:"3"` // this is the payer id
	Team      string      `protocp:"4"`
	Status    RoundStatus `protocp:"5"`
	State     RoundState  `protocp:"6"`
	Error     string      `protocp:"7"`
	CreatedAt time.Time   `protocp:"8"`
	UpdatedAt time.Time   `protocp:"9"`
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
	Players []RoundPlayerState `protocp:"1"`
}

type RoundPlayerState struct {
	Name      string         `protocp:"1"`
	Rank      int            `protocp:"2"`
	Parts     map[string]int32 `protocp:"3"`
	Included  bool           `protocp:"4"`
	Collected bool           `protocp:"5"`
	Submitted bool           `protocp:"6"`
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

func (rs RoundState) Included(player string)  bool {
	for _, p := range rs.Players {
		if p.Name != player {
			continue
		}

		return p.Included
	}

	return false
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

func LookupByIndex(ctx context.Context, dbc *sql.DB,index int) (*Round, error) {
	return lookupWhere(ctx,dbc,"index=?", index)
}

func Join(ctx context.Context, dbc *sql.DB, team string, matchID int64, index int) error {
	_, err := fsm.Insert(ctx, dbc, joinReq{Team: team, MatchID: matchID, Index: index})
	return err
}

func ToJoined(ctx context.Context, dbc *sql.DB, id int64, from RoundStatus,
	prevUpdatedAt time.Time, newState RoundState) error {

	return to(ctx, dbc, id, from, RoundStatusJoined, prevUpdatedAt,
		joinedReq{ID: id, State: newState})
}

func ToCollect(ctx context.Context, dbc *sql.DB, id int64, from RoundStatus,
	prevUpdatedAt time.Time, newState RoundState) error {

	return to(ctx, dbc, id, from, RoundStatusCollect, prevUpdatedAt,
		joinedReq{ID: id, State: newState})
}

func ToCollected(ctx context.Context, dbc *sql.DB, id int64, from RoundStatus,
	prevUpdatedAt time.Time, newState RoundState) error {

	return to(ctx, dbc, id, from, RoundStatusCollected, prevUpdatedAt,
		joinedReq{ID: id, State: newState})
}

func ensurePrevUpdatedAt(ctx context.Context, tx *sql.Tx, id int64, updatedAt time.Time) error {
	var n int
	err := tx.QueryRowContext(ctx, "select exists (select 1 from rounds "+
		"where id=? and updated_at=?)", id, updatedAt).Scan(&n)
	if err != nil {
		return err
	}

	if n != 1 {
		return engine.ErrConcurrentUpdates
	}

	return nil
}


func to(ctx context.Context, dbc *sql.DB, id int64, from, to RoundStatus,
	prevUpdatedAt time.Time, req shift.Updater) error {

	tx, err := dbc.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = ensurePrevUpdatedAt(ctx, tx, id, prevUpdatedAt)
	if err != nil {
		return err
	}

	notify, err := fsm.UpdateTx(ctx, tx, from, to, req)
	if err != nil {
		return err
	}
	defer notify()

	return tx.Commit()
}
