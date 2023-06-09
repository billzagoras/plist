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
	// for the time zones
	_, currentTimeLocationOffset := timeNowInLocation.Zone()

	// UTC t1
	timeObj1UTC, err := time.Parse("20060102T150405Z", t1)
	if utils.CheckErr(err) {
		return nil, errors.GetError(errors.TimeParsingError)
	}

	// UTC t2
	timeObj2UTC, err := time.Parse("20060102T150405Z", t2)
	if utils.CheckErr(err) {
		return nil, errors.GetError(errors.TimeParsingError)
	}

	errResp := utils.Round(&timeObj1UTC, &timeObj2UTC, period)
	if utils.CheckErr(errResp) {
		return nil, errResp
	}

	// Local t1
	timeObj1Local := timeObj1UTC.In(loc)
	_, timeObj1UTC = utils.NormalizeTime(timeObj1Local, timeObj1UTC)

	// Local t2
	timeObj2Local := timeObj2UTC.In(loc)
	_, timeObj2UTC = utils.NormalizeTime(timeObj2Local, timeObj2UTC)

	timestamps := []string{}
	for !timeObj1UTC.After(timeObj2UTC) {
		// Get the offset in seconds for each time zone for the time zones.
		_, givenTimeLocationOffset := timeObj1UTC.In(loc).Zone()
		duration := (currentTimeLocationOffset - givenTimeLocationOffset) / 60 / 60
		if duration != 0 {
			currentTimeLocationOffset = givenTimeLocationOffset
			timeObj1UTC = timeObj1UTC.Add(time.Duration(duration) * time.Hour)

			timeObj2UTC = timeObj2UTC.Add(time.Duration(duration) * time.Hour)
		}

		timestamps = append(timestamps, timeObj1UTC.Format("20060102T150405Z"))

		errResp := utils.AddPeriod(&timeObj1UTC, period)
		if utils.CheckErr(err) {
			return nil, errResp
		}
	}

	return &PtListResponse{
		Timestamps: timestamps,
	}, nil
}
