package humanreltime

import (
	"fmt"
	_ "log"
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

const MAXCOMPONENTS = 3

const (
	secYear   = 60 * 60 * 24 * 365
	secWeek   = 60 * 60 * 24 * 7
	secDay    = 60 * 60 * 24
	secHour   = 60 * 60
	secMinute = 60
)

type Prefix struct {
	Negative string
	Positive string
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
	Name      string
	LangCode  string
	Now       string
	Suffix    map[Resolution]Suffix
	Prefix    Prefix
	Delimiter Delimiter
}

func (l *Language) SuffixedNumber(number int64, res Resolution) string {
	switch {
	case number == 1 || number == -1:
		return fmt.Sprintf("%d %s", number, l.Suffix[res].Singular)
	case number == 0:
		return ""
	default:
		return fmt.Sprintf("%d %s", number, l.Suffix[res].Plural)
	}
}

func (l *Language) Duration(timestamp time.Time, reference time.Time, maxRes Resolution) string {
	var years, weeks, days, hours, minutes, seconds int64
	var allComponents = []*int64{&years, &weeks, &days, &hours, &minutes, &seconds}
	ref := reference.Unix()
	ts := timestamp.Unix()
	rest := ts - ref
	//log.Println(rest)
	// now
	if rest == 0 {
		return l.Now
	}
	// compute correct biggest unit
	switch maxRes {
	case Years:
		years = intAbs(rest / secYear)
		rest = rest % secYear
		//log.Println("[Y]", years, rest)
		fallthrough
	case Weeks:
		weeks = intAbs(rest / secWeek)
		rest = rest % secWeek
		//log.Println("[W]", weeks, rest)
		fallthrough
	case Days:
		days = intAbs(rest / secDay)
		rest = rest % secDay
		//log.Println("[D]", days, rest)
		fallthrough
	case Hours:
		hours = intAbs(rest / secHour)
		rest = rest % secHour
		fallthrough
	case Minutes:
		minutes = intAbs(rest / secMinute)
		rest = rest % secMinute
		fallthrough
	case Seconds:
		seconds = intAbs(rest)
	}
	//log.Println(years, weeks, days, hours, minutes, seconds)
	// remove zeros, generate strings
	components := make([]string, 0)
	for idx, res := range Resolutions {
		if *allComponents[idx] != 0 {
			components = append(components, l.SuffixedNumber(*allComponents[idx], res))
		}
	}
	// truncate resolution
	if len(components) > MAXCOMPONENTS {
		components = components[:MAXCOMPONENTS]
	}
	// handle special final delimiter
	if len(components) > 1 {
		front := components[:len(components)-2]
		tail := strings.Join(components[len(components)-2:len(components)], l.Delimiter.Last)
		components = append(front, tail)
	}
	timeString := strings.Join(components, l.Delimiter.Rest)
	// past or future?
	if ts-ref < 0 {
		return l.Prefix.Negative + timeString
	} else {
		return l.Prefix.Positive + timeString
	}
}

func intAbs(i int64) int64 {
	if i < 0 {
		return -1 * i
	} else {
		return i
	}
}
