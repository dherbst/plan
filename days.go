package plan

import "time"

// UntilResult contains the days between now and then.
type UntilResult struct {
	// Target is the date you are working to.
	Target string
	// Duration is the time.Duration between now and then.
	Duration time.Duration
	// Days is the number of calendar days.
	Days int
	// WorkingDays is the number of working days between now and then.
	WorkingDays int
	// Holidays is the number of Holidays between now and then.
	Holidays int
}

// TimeUntil returns the duration until the target date.
func TimeUntil(targetarg string) (time.Duration, error) {
	target, err := time.Parse("2006-01-02", targetarg)
	if err != nil {
		return 0, err
	}

	return time.Until(target), nil
}
