package main

import "time"

func yearDistinctCount(t time.Time) int {
	y, ok := years[t.Year()]
	var res int
	if ok {
		res = y.distinctCount
	}
	return res
}

func monthDistinctCount(t time.Time) int {
	y, ok := years[t.Year()]
	if !ok {
		return 0
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return 0
	}

	return m.distinctCount
}

func dayDistinctCount(t time.Time) int {
	y, ok := years[t.Year()]
	if !ok {
		return 0
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return 0
	}
	d := m.days[t.Day()-1]
	if d == nil {
		return 0
	}

	return d.distinctCount
}

func minuteDistinctCount(t time.Time) int {
	y, ok := years[t.Year()]
	if !ok {
		return 0
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return 0
	}
	d := m.days[t.Day()-1]
	if d == nil {
		return 0
	}
	h := d.hours[t.Hour()]
	if h == nil {
		return 0
	}

	mn := h.minutes[t.Minute()]
	if mn == nil {
		return 0
	}

	return mn.distinctCount
}

func secondDistinctCount(t time.Time) int {
	y, ok := years[t.Year()]
	if !ok {
		return 0
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return 0
	}
	d := m.days[t.Day()-1]
	if d == nil {
		return 0
	}
	h := d.hours[t.Hour()]
	if h == nil {
		return 0
	}

	mn := h.minutes[t.Minute()]
	if mn == nil {
		return 0
	}

	s := mn.seconds[t.Second()]
	if s == nil {
		return 0
	}

	return s.distinctCount
}
