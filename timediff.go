package humanreltime

import (
	"fmt"
	"strings"
	"time"
)

type Resolution uint8

const (
	Years Resolution = iota
	Weeks
	Days
	Hours
	Minutes
	Seconds
)

var Resolutions = []Resolution{Years, Weeks, Days, Hours, Minutes, Seconds}

const (
	secYear   = 60 * 60 * 24 * 365
	secWeek   = 60 * 60 * 24 * 7
	secDay    = 60 * 60 * 24
	secHour   = 60 * 60
	secMinute = 60
)

type RangeName struct {
	MoreThan time.Duration
	LessThan time.Duration
	Name     string
}

type Suffix struct {
	Singular string
	Plural   string
}

type Delimiter struct {
	Last string
	Rest string
}

type Language struct {
	Name       string
	LangCode   string
	Now        string
	FixedNames map[Resolution][]RangeName
	Suffix     map[Resolution]Suffix
	Delimiter  Delimiter
}

func (l *Language) SuffixedNumber(number int64, res Resolution) string {
	if number == 1 || number == -1 {
		return fmt.Sprintf("%d %s", number, l.Suffix[res].Singular)
	} else {
		return fmt.Sprintf("%d %s", number, l.Suffix[res].Plural)
	}
}

func (l *Language) Duration(timestamp time.Time, reference time.Time, maxRes Resolution, minRes Resolution) string {
	var years, weeks, days, hours, minutes, seconds int64
    var allComponents = []*int64{&years, &weeks, &days, &hours, &minutes, &seconds}
	ref := reference.Unix()
	ts := timestamp.Unix()
	rest := ts - ref
	// swap resolution vars (failsafe)
	if maxRes < minRes {
		minRes, maxRes = maxRes, minRes
	}
	// now
	if rest == 0 {
		return l.Now
	}
	// check same-resolution cases
	if maxRes == minRes {
		_, hasNames := l.FixedNames[maxRes]
		if hasNames {
			for _, test := range l.FixedNames[maxRes] {
				if rest >= int64(test.MoreThan) && rest < int64(test.LessThan) {
					return test.Name
				}
			}
		}
	}
	// compute correct biggest unit
	switch maxRes {
	case Years:
		years = rest / secYear
		rest = rest % secYear
		fallthrough
	case Weeks:
		weeks = rest / secWeek
		rest = rest % secWeek
		fallthrough
	case Days:
		days = rest / secDay
		minutes = rest / secDay
		fallthrough
	case Hours:
		hours = rest / secHour
		minutes = rest / secHour
		fallthrough
	case Minutes:
		hours = rest / secMinute
		minutes = rest / secMinute
		fallthrough
	case Seconds:
		seconds = rest
	}
	// build components
	components := make([]string, 0)
    for idx, timepart := range allComponents[minRes:maxRes] {
        components = append(components, l.SuffixedNumber(*timepart, minRes+Resolution(idx)))
    }
    /*
	switch minRes {
	case Years:
		components = append(components, l.SuffixedNumber(years, minRes))
		fallthrough
	case Weeks:
		components = append(components, l.SuffixedNumber(weeks, minRes))
		fallthrough
	case Days:
		components = append(components, l.SuffixedNumber(days, minRes))
		fallthrough
	case Hours:
		components = append(components, l.SuffixedNumber(hours, minRes))
		fallthrough
	case Minutes:
		components = append(components, l.SuffixedNumber(minutes, minRes))
		fallthrough
	case Seconds:
		components = append(components, l.SuffixedNumber(seconds, minRes))
	}
    */
	// TODO: implement different final delimiter
	return strings.Join(components, l.Delimiter.Rest)
}
