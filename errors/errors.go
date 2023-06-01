package errors

// Error codes.
const (
	UnsupportedPeriod    = 100
	TimeRoundingError    = 101
	TimezoneLoadingError = 102
	TimeParsingError     = 103
	AddingPeriodError    = 104
)

// Error struct.
type ErrResp struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

// Assign error codes to errResp struct.
var ErrorMap = map[int]ErrResp{
	UnsupportedPeriod: {
		Status: "error",
		Desc:   "Failed to round time objects",
	},
	TimeRoundingError: {
		Status: "error",
		Desc:   "Failed to round time objects",
	},
	TimezoneLoadingError: {
		Status: "error",
		Desc:   "Could not load given timezone location",
	},
	TimeParsingError: {
		Status: "error",
		Desc:   "Could not parse time in go",
	},
	AddingPeriodError: {
		Status: "error",
		Desc:   "Could not add period to time object",
	},
}

// Retrieve a new error object.
func GetError(errorCode int) *ErrResp {
	resp := &ErrResp{
		Status: ErrorMap[errorCode].Status,
		Desc:   ErrorMap[errorCode].Desc,
	}
	return resp
}
