package humanreltime

import (
	"testing"
	"time"
)

func TestGerman(t *testing.T) {
	now := time.Now()
	tests := []struct {
		diff   time.Duration
		maxRes Resolution
		expect string
	}{
		{0, Years, "jetzt"},
		{1 * time.Second, Years, "in 1 Sekunde"},
		{23 * time.Second, Years, "in 23 Sekunden"},
		{-1 * time.Second, Years, "vor 1 Sekunde"},
		{-42 * time.Second, Years, "vor 42 Sekunden"},
		{2*time.Minute + 5*time.Second, Years, "in 2 Minuten und 5 Sekunden"},
		{1*time.Second + 2*time.Minute + 3*time.Hour + 4*secDay*time.Second + 5*secWeek*time.Second + 6*secYear*time.Second, Years, "in 6 Jahren, 5 Wochen und 4 Tagen"},
		{1*time.Second + 2*time.Minute + 3*time.Hour + 4*secDay*time.Second + 5*secWeek*time.Second + 6*secYear*time.Second, Weeks, "in 318 Wochen, 3 Tagen und 3 Stunden"},
	}
	for num, test := range tests {
		s := German.Duration(now.Add(test.diff), now, test.maxRes)
		if s != test.expect {
			t.Errorf("%02d: Got \"%s\", expected \"%s\"", num, s, test.expect)
		}
	}
}
