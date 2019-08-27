package main

import "testing"

func TestPopulateDB(t *testing.T) {
	populateDB("hn_logs_test.tsv")

	if len(records) != 10 {
		t.Errorf("have %d records. want %d\n", len(records), 10)
	}

	l := "kevingosse.com"
	u, ok := urls[l]
	if !ok {
		t.Errorf("expect %s to be in urls\n", l)
	}

	if iurls[u] != l {
		t.Errorf("have iurls[%d]=%s. want %s\n", u, iurls[u], l)
	}

	yearsTestCases := []struct {
		value    int
		expected bool
	}{
		{2014, false},
		{2015, true},
		{2016, true},
	}

	for _, tc := range yearsTestCases {
		_, ok := years[tc.value]
		if ok != tc.expected {
			t.Errorf("have years[%d] = _, %v. want _, %v\n", tc.value, ok, tc.expected)
		}
	}

	// 2016-01-01 00:03:50
	y := years[2016]
	if y == nil {
		t.Error("should not have nil year")
		t.FailNow()
	}
	m := y.months[0]
	if m == nil {
		t.Error("should not have nil month")
		t.FailNow()
	}
	d := m.days[0]
	if d == nil {
		t.Error("should not have nil day")
		t.FailNow()
	}
	h := d.hours[0]
	if h == nil {
		t.Error("should not have nil hour")
		t.FailNow()
	}
	mn := h.minutes[3]
	if mn == nil {
		t.Error("should not have nil minute")
		t.FailNow()
	}
	s := mn.seconds[50]
	if s == nil {
		t.Error("should not have nil second")
	}

}
