package main

func updateYear(y *year, r *record) error {

	updateDisctinctCount(&y.distinctCount, y.urls, &y.sortedurls, r.url)

	i := int(r.time.Month()) - 1
	m := y.months[i]
	if m == nil {
		m = &month{
			id:   i + 1,
			urls: make(map[string]int),
		}
		y.months[i] = m
	}

	return updateMonth(m, r)
}

func updateMonth(m *month, r *record) error {

	updateDisctinctCount(&m.distinctCount, m.urls, &m.sortedurls, r.url)

	i := r.time.Day() - 1
	d := m.days[i]
	if d == nil {
		d = &day{
			id:   i + 1,
			urls: make(map[string]int),
		}
		m.days[i] = d
	}

	return udpateDay(d, r)
}

func udpateDay(d *day, r *record) error {

	updateDisctinctCount(&d.distinctCount, d.urls, &d.sortedurls, r.url)

	i := r.time.Hour()
	h := d.hours[i]
	if h == nil {
		h = &hour{
			id:   i,
			urls: make(map[string]int),
		}
		d.hours[i] = h
	}
	return updateHour(h, r)
}

func updateHour(h *hour, r *record) error {

	updateDisctinctCount(&h.distinctCount, h.urls, &h.sortedurls, r.url)

	i := r.time.Minute()
	m := h.minutes[i]
	if m == nil {
		m = &minute{
			id:   i,
			urls: make(map[string]int),
		}
		h.minutes[i] = m
	}
	return updateMinute(m, r)
}

func updateMinute(m *minute, r *record) error {

	updateDisctinctCount(&m.distinctCount, m.urls, &m.sortedurls, r.url)

	i := r.time.Second()
	s := m.seconds[i]
	if s == nil {
		s = &second{
			id:   i,
			urls: make(map[string]int),
		}
		m.seconds[i] = s
	}
	return updateSecond(s, r)
}

func updateSecond(s *second, r *record) error {

	updateDisctinctCount(&s.distinctCount, s.urls, &s.sortedurls, r.url)

	s.records = append(s.records, r.id)

	return nil
}
