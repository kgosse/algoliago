package main

func updateYear(y *year, r *record) error {

	v, ok := y.urls[r.url]
	y.urls[r.url] = v + 1
	if !ok {
		y.distinctCount++
		y.sortedurls = append(y.sortedurls, urls[r.url])
	}

	i := int(r.time.Month()) - 1
	m := y.months[i]
	if m == nil {
		m = &month{
			id:   i,
			urls: make(map[string]int),
		}
		y.months[i] = m
	}

	return updateMonth(m, r)
}

func updateMonth(m *month, r *record) error {
	v, ok := m.urls[r.url]
	m.urls[r.url] = v + 1
	if !ok {
		m.distinctCount++
		m.sortedurls = append(m.sortedurls, urls[r.url])
	}

	i := r.time.Day()
	d := m.days[i]
	if d == nil {
		d = &day{
			id:   i,
			urls: make(map[string]int),
		}
		m.days[i] = d
	}

	return udpateDay(d, r)
}

func udpateDay(d *day, r *record) error {
	v, ok := d.urls[r.url]
	d.urls[r.url] = v + 1
	if !ok {
		d.distinctCount++
		d.sortedurls = append(d.sortedurls, urls[r.url])
	}

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
	v, ok := h.urls[r.url]
	h.urls[r.url] = v + 1
	if !ok {
		h.distinctCount++
		h.sortedurls = append(h.sortedurls, urls[r.url])
	}

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
	v, ok := m.urls[r.url]
	m.urls[r.url] = v + 1
	if !ok {
		m.distinctCount++
		m.sortedurls = append(m.sortedurls, urls[r.url])
	}

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
	v, ok := s.urls[r.url]
	s.urls[r.url] = v + 1
	if !ok {
		s.distinctCount++
		s.sortedurls = append(s.sortedurls, urls[r.url])
	}

	s.records = append(s.records, r.id)

	return nil
}
