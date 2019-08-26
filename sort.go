package main

import "sort"

func sortYearUrls(y *year) {
	if y == nil {
		return
	}
	sort.Slice(y.sortedurls, func(i, j int) bool {
		iv, _ := y.urls[iurls[y.sortedurls[i]]]
		jv, _ := y.urls[iurls[y.sortedurls[j]]]
		return iv > jv
	})
	for i := 0; i < len(y.months); i++ {
		sortMonthUrls(y.months[i])
	}
}

func sortMonthUrls(m *month) {
	if m == nil {
		return
	}
	sort.Slice(m.sortedurls, func(i, j int) bool {
		iv, _ := m.urls[iurls[m.sortedurls[i]]]
		jv, _ := m.urls[iurls[m.sortedurls[j]]]
		return iv > jv
	})
	for i := 0; i < len(m.days); i++ {
		sortDayUrls(m.days[i])
	}
}

func sortDayUrls(d *day) {
	if d == nil {
		return
	}
	sort.Slice(d.sortedurls, func(i, j int) bool {
		iv, _ := d.urls[iurls[d.sortedurls[i]]]
		jv, _ := d.urls[iurls[d.sortedurls[j]]]
		return iv > jv
	})
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
	sort.Slice(m.sortedurls, func(i, j int) bool {
		iv, _ := m.urls[iurls[m.sortedurls[i]]]
		jv, _ := m.urls[iurls[m.sortedurls[j]]]
		return iv > jv
	})
	for i := 0; i < len(m.seconds); i++ {
		sortSecondUrls(m.seconds[i])
	}
}

func sortSecondUrls(s *second) {
	if s == nil {
		return
	}
	sort.Slice(s.sortedurls, func(i, j int) bool {
		iv, _ := s.urls[iurls[s.sortedurls[i]]]
		jv, _ := s.urls[iurls[s.sortedurls[j]]]
		return iv > jv
	})
}
