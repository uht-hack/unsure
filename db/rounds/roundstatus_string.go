// Code generated by "stringer -type=RoundStatus -trimprefix=RoundStatus"; DO NOT EDIT.

package rounds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RoundStatusUnknown-0]
	_ = x[RoundStatusJoin-1]
	_ = x[RoundStatusJoined-2]
	_ = x[RoundStatusCollect-3]
	_ = x[RoundStatusCollected-4]
	_ = x[RoundStatusSubmit-5]
	_ = x[RoundStatusSubmitted-6]
	_ = x[RoundStatusSuccess-7]
	_ = x[RoundStatusFailed-8]
	_ = x[roundStatusSentinel-9]
}

const _RoundStatus_name = "UnknownJoinJoinedCollectCollectedSubmitSubmittedSuccessFailedroundStatusSentinel"

var _RoundStatus_index = [...]uint8{0, 7, 11, 17, 24, 33, 39, 48, 55, 61, 80}

func (i RoundStatus) String() string {
	if i < 0 || i >= RoundStatus(len(_RoundStatus_index)-1) {
		return "RoundStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RoundStatus_name[_RoundStatus_index[i]:_RoundStatus_index[i+1]]
}