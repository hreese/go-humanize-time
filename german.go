package humanreltime

var German = Language{
	Name:     "Deutsch",
	LangCode: "de",
	Now:      "jetzt",
	Prefix: Prefix{
		Negative: "vor ",
		Positive: "in ",
	},
	Suffix: map[Resolution]Suffix{
		Years: Suffix{
			Singular: "Jahr",
			Plural:   "Jahren",
		},
		Months: Suffix{
			Singular: "Monat",
			Plural:   "Monaten",
		},
		Weeks: Suffix{
			Singular: "Woche",
			Plural:   "Wochen",
		},
		Days: Suffix{
			Singular: "Tag",
			Plural:   "Tagen",
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
		Last: " und ",
		Rest: ", ",
	},
}
