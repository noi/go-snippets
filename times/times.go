package times

import "time"

type AddMode uint8

const (
	// This mode works the same as `time.Time.AddDate`.
	// e.g. 3/31 + 1 month is 5/1
	NormalizeExcessDays AddMode = iota
	// This mode truncates excess days without normalizing.
	// e.g. 3/31 + 1 month is 4/30
	TruncateExcessDays
	// This mode is basically the same as `TruncateExcessDays`,
	// but if the original time is the end of the month,
	// the result will also be the end of the month.
	// e.g. 3/31 + 1 month is 4/30, 4/30 + 1 month is 5/31
	PreserveEndOfMonth
)

var JST *time.Location

func BeginDayOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, -t.Day()+1)
}

func EndDayOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, -t.Day())
}

func TruncateToDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func AddYears(t time.Time, years int, mode AddMode) time.Time {
	return addYearsAndMonths(t, years, 0, mode)
}

func AddMonths(t time.Time, months int, mode AddMode) time.Time {
	return addYearsAndMonths(t, 0, months, mode)
}

func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func addYearsAndMonths(t time.Time, years, months int, mode AddMode) time.Time {
	if mode == NormalizeExcessDays {
		return t.AddDate(years, months, 0)
	}
	day := t.Day()
	year, month, dayLimit := t.AddDate(years, months+1, -t.Day()).Date()
	if day > dayLimit || mode == PreserveEndOfMonth && t.Month() != t.AddDate(0, 0, 1).Month() {
		day = dayLimit
	}
	hour, min, sec := t.Clock()
	return time.Date(year, month, day, hour, min, sec, t.Nanosecond(), t.Location())
}

func init() {
	if loc, err := time.LoadLocation("Asia/Tokyo"); err == nil {
		JST = loc
	} else {
		JST = time.FixedZone("JST", 9*60*60)
	}
}
