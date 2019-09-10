package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPopularHandler(t *testing.T) {
	populateDB("hn_logs.tsv")
	ts := httptest.NewServer(newHandler())
	defer ts.Close()

	testCases := []struct {
		title string // title of the test
		route string // the route being tested
		out   string // the expected result
	}{
		{
			"route /1/queries/popular/2015?size=3",
			"/1/queries/popular/2015?size=3",
			`{
				"queries": [
				  { "query": "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice", "count": 6675 },
				  { "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045", "count": 4652 },
				  { "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1", "count": 3100 }
				]
			  }`,
		},
		{
			"route /1/queries/popular/2015-08-02?size=5",
			"/1/queries/popular/2015-08-02?size=5",
			`{
				"queries": [
					{ "query": "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice", "count": 2283 },
					{ "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045", "count": 1943 },
					{ "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1", "count": 1358 },
					{ "query": "http%3A%2F%2Fjamonkey.com%2F50-organizing-ideas-for-every-room-in-your-house%2F", "count": 890 },
					{ "query": "http%3A%2F%2Fsharingis.cool%2F1000-musicians-played-foo-fighters-learn-to-fly-and-it-was-epic", "count": 701 }
				]
			  }`,
		},
		{
			"route /1/queries/popular/2015-08-01%2000:03:50?size=2",
			"/1/queries/popular/2015-08-01%2000:03:50?size=2",
			`{
				"queries": [
				  { "query": "%22http%3A%2F%2Fnewpickuptrucks2016.com%2F2016-chevrolet-el-camino-coming-back%2F%22", "count": 1 },
				  { "query": "%22http%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DaBKRgTNqc5E%22", "count": 1 }
				]
			  }`,
		},
		{
			"route /1/queries/popular/2015-08?size=3",
			"/1/queries/popular/2015-08?size=3",
			`{
				"queries": [
				  { "query": "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice", "count": 6675 },
				  { "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045", "count": 4652 },
				  { "query": "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1", "count": 3100 }
				]
			  }`,
		},
		{
			"route /1/queries/popular/2015-08-04%2005:57?size=1",
			"/1/queries/popular/2015-08-04%2005:57?size=1",
			`{
				"queries": [
				  { "query": "docker", "count": 2 }
				]
			  }`,
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			res, err := http.Get(ts.URL + test.route)
			if err != nil {
				t.Error(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Error(err)
			}

			var testJSON popularResponse
			err = json.Unmarshal([]byte(test.out), &testJSON)
			if err != nil {
				t.Errorf("Failed to json.Unmarshal %s\n", test.out)
				t.FailNow()
			}

			var bodyJSON popularResponse
			err = json.Unmarshal(body, &bodyJSON)
			if err != nil {
				t.Errorf("Failed to json.Unmarshal %s\n", body)
				t.FailNow()
			}
			if !reflect.DeepEqual(testJSON, bodyJSON) {
				t.Errorf("have %s\nwant %s", body, test.out)
			}
		})
	}
}

func TestCountHandler(t *testing.T) {
	populateDB("hn_logs.tsv")
	ts := httptest.NewServer(newHandler())
	defer ts.Close()

	testCases := []struct {
		title string // title of the test
		route string // the route being tested
		out   string // the expected result
	}{
		{
			"route /1/queries/count/2015",
			"/1/queries/count/2015",
			"{\"count\":573697}",
		},
		{
			"route /1/queries/count/2015-08",
			"/1/queries/count/2015-08",
			"{\"count\":573697}",
		},
		{
			"route /1/queries/count/2015-08-03",
			"/1/queries/count/2015-08-03",
			"{\"count\":198117}",
		},
		{
			"route /1/queries/count/2015-08-01%2000:04",
			"/1/queries/count/2015-08-01%2000:04",
			"{\"count\":617}",
		},
		{
			"route /1/queries/count/2015-08-01%2000:03:50",
			"/1/queries/count/2015-08-01%2000:03:50",
			"{\"count\":10}",
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			res, err := http.Get(ts.URL + test.route)
			if err != nil {
				t.Error(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Error(err)
			}

			var testJSON countResponse
			err = json.Unmarshal([]byte(test.out), &testJSON)
			if err != nil {
				t.Errorf("Failed to json.Unmarshal %s\n", test.out)
				t.FailNow()
			}

			var bodyJSON countResponse
			err = json.Unmarshal(body, &bodyJSON)
			if err != nil {
				t.Errorf("Failed to json.Unmarshal %s\n", body)
				t.FailNow()
			}
			if !reflect.DeepEqual(testJSON, bodyJSON) {
				t.Errorf("have %s\nwant %s", body, test.out)
			}
		})
	}
}
