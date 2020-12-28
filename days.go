package plan

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
)

// UntilResult contains the days between now and then.
type UntilResult struct {
	// Target is the date you are working to.
	Target time.Time
	// Duration is the time.Duration between now and then.
	Duration time.Duration
	// Days is the number of calendar days.
	Days int
	// WorkingDays is the number of working days between now and then.
	WorkingDays int
	// Holidays is the number of Holidays between now and then.
	Holidays int
}

// String is a formatted UntilResult.
func (r *UntilResult) String() string {
	value := fmt.Sprintf("Days: %v\nWorkingDays: %v\nHolidays: %v\n", r.Days, r.WorkingDays, r.Holidays)
	return value
}

// TimeUntil returns the duration until the target date.
func TimeUntil(targetarg string) (*UntilResult, error) {

	target, err := time.Parse("2006-01-02", targetarg)
	if err != nil {
		return nil, err
	}

	c := cal.NewBusinessCalendar()
	DayAfterThanksgiving := &cal.Holiday{
		Name:    "Day After Thanksgiving",
		Type:    cal.ObservancePublic,
		Month:   time.November,
		Weekday: time.Friday,
		Offset:  5,
		Func:    cal.CalcWeekdayOffset,
	}
	c.AddHoliday(
		us.NewYear,
		us.MlkDay,
		us.PresidentsDay,
		us.MemorialDay,
		us.IndependenceDay,
		us.LaborDay,
		us.ThanksgivingDay,
		DayAfterThanksgiving,
		us.ChristmasDay,
	)

	duration := time.Until(target)
	result := &UntilResult{
		Target:   target,
		Duration: duration,
		Days:     int(duration.Hours() / 24),
	}

	for today := time.Now(); today.Before(target); today = today.AddDate(0, 0, 1) {
		_, observedHoliday, _ := c.IsHoliday(today)
		if observedHoliday {
			result.Holidays++
		} else if c.IsWorkday(today) {
			result.WorkingDays++
		}
	}

	return result, nil
}
