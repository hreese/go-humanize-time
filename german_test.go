package humanreltime

import (
	"testing"
	"time"
)

func TestGerman(t *testing.T) {
	now := time.Now()
	tests := []struct {
		diff          time.Duration
		maxRes        Resolution
		NumComponents int
		expect        string
	}{
		{0, Years, 2, "jetzt"},
		{1 * time.Second, Years, 2, "in 1 Sekunde"},
		{23 * time.Second, Years, 2, "in 23 Sekunden"},
		{-1 * time.Second, Years, 2, "vor 1 Sekunde"},
		{-42 * time.Second, Years, 2, "vor 42 Sekunden"},
		{2*time.Minute + 5*time.Second, Years, 2, "in 2 Minuten und 5 Sekunden"},
		{1*time.Second + 2*time.Minute + 3*time.Hour + 4*secDay*time.Second + 5*secWeek*time.Second + 6*secYear*time.Second, Years, 2, "in 6 Jahren und 1 Monat"},
		{1*time.Second + 2*time.Minute + 3*time.Hour + 4*secDay*time.Second + 5*secWeek*time.Second + 6*secYear*time.Second, Months, 2, "in 74 Monaten und 1 Woche"},
	}
	for num, test := range tests {
		s := German.Duration(now.Add(test.diff), now, test.maxRes, test.NumComponents)
		if s != test.expect {
			t.Errorf("%02d: Got \"%s\", expected \"%s\"", num, s, test.expect)
		}
	}
}
