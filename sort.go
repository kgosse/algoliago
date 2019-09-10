package main

func sortYearUrls(y *year) {
	if y == nil {
		return
	}
	sortURL(y.sortedurls, y.urls)
	for i := 0; i < len(y.months); i++ {
		sortMonthUrls(y.months[i])
	}
}

func sortMonthUrls(m *month) {
	if m == nil {
		return
	}
	sortURL(m.sortedurls, m.urls)
	for i := 0; i < len(m.days); i++ {
		sortDayUrls(m.days[i])
	}
}

func sortDayUrls(d *day) {
	if d == nil {
		return
	}
	sortURL(d.sortedurls, d.urls)
	for i := 0; i < len(d.hours); i++ {
		sortHourUrls(d.hours[i])
	}
}

func sortHourUrls(h *hour) {
	if h == nil {
		return
	}

	for i := 0; i < len(h.minutes); i++ {
		sortMinuteUrls(h.minutes[i])
	}
}

func sortMinuteUrls(m *minute) {
	if m == nil {
		return
	}
	sortURL(m.sortedurls, m.urls)
	for i := 0; i < len(m.seconds); i++ {
		sortSecondUrls(m.seconds[i])
	}
}

func sortSecondUrls(s *second) {
	if s == nil {
		return
	}
	sortURL(s.sortedurls, s.urls)
}
