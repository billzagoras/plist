package ptlist

import (
	"context"
	"plist/errors"
	"plist/utils"
	"time"
)

// Service struct represents ptlist service.
type Service struct{}

// NewService service constructor.
func NewService() *Service {
	return &Service{}
}

// GetPtList returns a list of all matching timestamps of a periodic task between 2 time points
// in UTC in the following form: 20060102T150405Z.
func (s *Service) GetPtList(ctx context.Context, period, tz, t1, t2 string) (*PtListResponse, *errors.ErrResp) {
	loc, err := time.LoadLocation(tz)
	if utils.CheckErr(err) {
		return nil, errors.GetError(errors.TimezoneLoadingError)
	}

	// Get the current time in UTC and the specified location.
	timeNowInLocation := time.Now().In(loc)

	// Get the offset in seconds for each time zone
	//for the time zones
	_, currentTimeLocationOffset := timeNowInLocation.Zone()

	timeObj1, err := time.ParseInLocation("20060102T150405Z", t1, loc)
	if utils.CheckErr(err) {
		return nil, errors.GetError(errors.TimeParsingError)
	}

	timeObj2, err := time.ParseInLocation("20060102T150405Z", t2, loc)
	if utils.CheckErr(err) {
		return nil, errors.GetError(errors.TimeParsingError)
	}

	errResp := utils.Round(&timeObj1, &timeObj2, period)
	if utils.CheckErr(errResp) {
		return nil, errResp
	}

	timestamps := []string{}
	for !timeObj1.After(timeObj2) {
		// Get the offset in seconds for each time zone for the time zones.
		_, givenTimeLocationOffset := timeObj1.Zone()
		duration := (currentTimeLocationOffset - givenTimeLocationOffset) / 60 / 60
		if duration != 0 {
			currentTimeLocationOffset = givenTimeLocationOffset
			timeObj1 = timeObj1.Add(time.Duration(duration) * time.Hour)

			// Monkey business :p
			timeObj2 = timeObj2.Add(time.Duration(duration) * time.Hour)
		}

		timestamps = append(timestamps, timeObj1.Format("20060102T150405Z"))

		errResp := utils.AddPeriod(&timeObj1, period)
		if utils.CheckErr(err) {
			return nil, errResp
		}
	}

	return &PtListResponse{
		Timestamps: timestamps,
	}, nil
}
