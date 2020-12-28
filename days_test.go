package plan

import (
	"testing"
	"time"
)

func TestTimeUntil(t *testing.T) {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	t.Log(tomorrow.Format("2006-01-02"))
	duration, err := TimeUntil(tomorrow.Format("2006-01-02"))
	if err != nil {
		t.Fatalf("Error TimeUntil=%v\n", err)
	}
	if duration.Seconds() < 1 {
		t.Fatalf("Didn't expect to be less than a second")
	}
}
