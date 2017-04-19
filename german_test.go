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
		minRes Resolution
		expect string
	}{
		{0, Years, Seconds, "jetzt"},
		{time.Second, Years, Seconds, "1 Sekunde"},
	}
	for num, test := range tests {
		s := German.Duration(now.Add(test.diff), now, test.maxRes, test.minRes)
		if s != test.expect {
			t.Errorf("%02d: Got \"%s\", expected \"%s\"", num, s, test.expect)
		}
	}
}
