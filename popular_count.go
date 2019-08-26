package main

import "time"

func yearPopularCount(t time.Time, s int) popularResponse {
	y, ok := years[t.Year()]
	var res popularResponse
	if ok {
		for i := 0; i < s; i++ {
			u := iurls[y.sortedurls[i]]
			v := y.urls[u]
			q := queryItem{
				u,
				v,
			}
			res.Queries = append(res.Queries, q)
		}
	}
	return res
}

func monthPopularCount(t time.Time, s int) (res popularResponse) {
	y, ok := years[t.Year()]
	if !ok {
		return
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return
	}
	for i := 0; i < s; i++ {
		u := iurls[m.sortedurls[i]]
		v := m.urls[u]
		q := queryItem{
			u,
			v,
		}
		res.Queries = append(res.Queries, q)
	}
	return
}

func dayPopularCount(t time.Time, s int) (res popularResponse) {
	y, ok := years[t.Year()]
	if !ok {
		return
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return
	}
	d := m.days[t.Day()]
	if d == nil {
		return
	}
	for i := 0; i < s; i++ {
		u := iurls[d.sortedurls[i]]
		v := d.urls[u]
		q := queryItem{
			u,
			v,
		}
		res.Queries = append(res.Queries, q)
	}
	return
}

func minutePopularCount(t time.Time, s int) (res popularResponse) {
	y, ok := years[t.Year()]
	if !ok {
		return
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return
	}
	d := m.days[t.Day()]
	if d == nil {
		return
	}
	h := d.hours[t.Hour()]
	if h == nil {
		return
	}
	mn := h.minutes[t.Minute()]
	if mn == nil {
		return
	}

	for i := 0; i < s; i++ {
		u := iurls[mn.sortedurls[i]]
		v := d.urls[u]
		q := queryItem{
			u,
			v,
		}
		res.Queries = append(res.Queries, q)
	}
	return
}

func secondPopularCount(t time.Time, s int) (res popularResponse) {
	y, ok := years[t.Year()]
	if !ok {
		return
	}
	m := y.months[int(t.Month())-1]
	if m == nil {
		return
	}
	d := m.days[t.Day()]
	if d == nil {
		return
	}
	h := d.hours[t.Hour()]
	if h == nil {
		return
	}
	mn := h.minutes[t.Minute()]
	if mn == nil {
		return
	}

	sec := mn.seconds[t.Second()]
	if sec == nil {
		return
	}

	for i := 0; i < s; i++ {
		u := iurls[sec.sortedurls[i]]
		v := d.urls[u]
		q := queryItem{
			u,
			v,
		}
		res.Queries = append(res.Queries, q)
	}
	return
}
