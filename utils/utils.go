package utils

import (
	"log"
	"plist/errors"
	"reflect"
	"time"
)

// Round time objects to closest time objects based on the requirements.
func Round(timeObj1, timeObj2 *time.Time, period string) *errors.ErrResp {
	if timeObj1.Minute() > 0 {
		*timeObj1 = time.Date(timeObj1.Year(), timeObj1.Month(), timeObj1.Day(), timeObj1.Add(time.Hour).Hour(), 0, 0, 0, timeObj1.Location())
	}

	switch period {
	case "1h":
		if timeObj2.Minute() > 0 {
			*timeObj2 = time.Date(timeObj2.Year(), timeObj2.Month(), timeObj2.Day(), timeObj2.Hour(), 0, 0, 0, timeObj2.Location())
		}
		return nil
	case "1d":
		if timeObj2.Minute() > 0 {
			*timeObj2 = time.Date(timeObj2.Year(), timeObj2.Month(), timeObj2.Day()-1, timeObj1.Hour(), 0, 0, 0, timeObj2.Location())
		}
		return nil
	case "1mo":
		// Last day of the month trick :)
		*timeObj1 = time.Date(timeObj1.Year(), timeObj1.Month()+1, 0, timeObj1.Hour(), 0, 0, 0, timeObj1.Location())

		if timeObj2.Day() > 0 {
			*timeObj2 = time.Date(timeObj2.Year(), timeObj2.Month(), 0, timeObj1.Hour(), 0, 0, 0, timeObj2.Location())
		}

		return nil
	case "1y":
		// Last month and day of the month trick :)
		*timeObj1 = time.Date(timeObj1.Year()+1, 1, 0, timeObj1.Hour(), 0, 0, 0, timeObj1.Location())

		if timeObj2.Month() > 0 {
			*timeObj2 = time.Date(timeObj2.Year(), 1, 0, timeObj1.Hour(), 0, 0, 0, timeObj2.Location())
		}
		return nil
	}

	return errors.GetError(errors.TimeRoundingError)
}

// Add periods of time to time object until that reaches the end of given time.
func AddPeriod(timeObj *time.Time, period string) *errors.ErrResp {
	switch period {
	case "1h":
		*timeObj = timeObj.Add(time.Hour)
		return nil
	case "1d":
		*timeObj = timeObj.AddDate(0, 0, 1)
		return nil
	case "1mo":
		*timeObj = time.Date(timeObj.Year(), timeObj.Month()+2, 0, timeObj.Hour(), 0, 0, 0, timeObj.Location())
		return nil
	case "1y":
		*timeObj = time.Date(timeObj.Year()+1, 1, 0, timeObj.Hour(), 0, 0, 0, timeObj.Location())
		*timeObj = timeObj.AddDate(1, 0, 0)
		return nil
	}

	return errors.GetError(errors.AddingPeriodError)
}

func NormalizeTime(timeObjLocal, timeObjUTC time.Time) (time.Time, time.Time) {
	if timeObjLocal.Format("20060102") > timeObjUTC.Format("20060102") {
		timeObjLocal = time.Date(timeObjLocal.Year(), timeObjLocal.Month(), timeObjLocal.Day(), 0, 0, 0, 0, timeObjLocal.Location())
		timeObjUTC = timeObjLocal.UTC()
	} else if timeObjLocal.Format("20060102") < timeObjUTC.Format("20060102") {
		timeObjLocal = time.Date(timeObjLocal.Year(), timeObjLocal.Month(), timeObjLocal.Day(), 24, 0, 0, 0, timeObjLocal.Location())
		timeObjUTC = timeObjLocal.UTC()
	}
	return timeObjLocal, timeObjUTC
}

// Checks for errors.
func CheckErr(err interface{}) bool {
	if err != nil && !reflect.ValueOf(err).IsNil() {
		errorType := reflect.TypeOf(err).Elem().String()
		switch errorType {
		case "errors.errorString":
			log.Println(err)
			return true
		case "errors.ErrResp":
			log.Println(err)
			return true
		}
	}
	return false
}
