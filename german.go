package humanreltime

import (
	"time"
)

var German = Language{
	Name:     "Deutsch",
	LangCode: "de",
	Now:      "jetzt",
	FixedNames: map[Resolution][]RangeName{
		Years: []RangeName{
			RangeName{
				MoreThan: -1 * time.Hour * 24 * 365 * 2,
				LessThan: -1 * time.Hour * 24 * 365,
				Name:     "vorletztes Jahr",
			},
			RangeName{
				MoreThan: -1 * time.Hour * 24 * 365,
				LessThan: 0,
				Name:     "letztes Jahr",
			},
			RangeName{
				MoreThan: 0,
				LessThan: time.Hour * 24 * 365,
				Name:     "dieses Jahr",
			},
			RangeName{
				MoreThan: time.Hour * 24 * 365,
				LessThan: time.Hour * 24 * 365 * 2,
				Name:     "nächstes Jahr",
			},
			RangeName{
				MoreThan: time.Hour * 24 * 365 * 2,
				LessThan: time.Hour * 24 * 365 * 3,
				Name:     "übernächstes Jahr",
			},
		},
		Weeks: []RangeName{
			RangeName{
				MoreThan: -1 * time.Hour * 24 * 7 * 2,
				LessThan: -1 * time.Hour * 24 * 7,
				Name:     "vorletzte Woche",
			},
			RangeName{
				MoreThan: -1 * time.Hour * 24 * 7,
				LessThan: 0,
				Name:     "letzte Woche",
			},
			RangeName{
				MoreThan: 0,
				LessThan: time.Hour * 24 * 7,
				Name:     "diese Woche",
			},
			RangeName{
				MoreThan: time.Hour * 24 * 7,
				LessThan: time.Hour * 24 * 7 * 2,
				Name:     "nächste Woche",
			},
			RangeName{
				MoreThan: time.Hour * 24 * 7 * 2,
				LessThan: time.Hour * 24 * 7 * 3,
				Name:     "übernächste Woche",
			},
		},
		Days: []RangeName{
			RangeName{
				MoreThan: -1 * time.Hour * 24 * 2,
				LessThan: -1 * time.Hour * 24,
				Name:     "vorgestern",
			},
			RangeName{
				MoreThan: -1 * time.Hour * 24,
				LessThan: 0,
				Name:     "gestern",
			},
			RangeName{
				MoreThan: 0,
				LessThan: time.Hour * 24,
				Name:     "heute",
			},
			RangeName{
				MoreThan: time.Hour * 24,
				LessThan: time.Hour * 24 * 2,
				Name:     "morgen",
			},
			RangeName{
				MoreThan: time.Hour * 24 * 2,
				LessThan: time.Hour * 24 * 3,
				Name:     "übermorgen",
			},
		},
	},
	Suffix: map[Resolution]Suffix{
		Years: Suffix{
			Singular: "Jahr",
			Plural:   "Jahre",
		},
		Weeks: Suffix{
			Singular: "Woche",
			Plural:   "Wochen",
		},
		Hours: Suffix{
			Singular: "Stunde",
			Plural:   "Stunden",
		},
		Minutes: Suffix{
			Singular: "Minute",
			Plural:   "Minuten",
		},
		Seconds: Suffix{
			Singular: "Sekunde",
			Plural:   "Sekunden",
		},
	},
	Delimiter: Delimiter{
		Last: "und",
		Rest: ",",
	},
}
