package plan

import "time"

// TimeUntil returns the duration until the target date.
func TimeUntil(targetarg string) (time.Duration, error) {
	target, err := time.Parse("2006-01-02", targetarg)
	if err != nil {
		return 0, err
	}

	return time.Until(target), nil
}
