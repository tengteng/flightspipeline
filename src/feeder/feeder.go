package feeder

// https://developer.flightstats.com/api-docs/flight-status-feed/v2/flight-status-feed-api
// app id is 69080051 and our app key is 00cacbefbb77f944e59dd1e72b132cdd

import (
	"net/http"

	"client"

	"github.com/golang/glog"
)

type Resp struct {
	Request map[string]string `json:"request"`
	Item    string            `json:"item"`
}

type FlightStatsClient struct {
	client http.Client
}

const (
	FlightStatsURL = "http://api.flightstats.com/flex/flightstatusfeed/rest/v2/json/latest"
	AppID          = "69080051"
	AppKey         = "00cacbefbb77f944e59dd1e72b132cdd"
)

func (c *FlightStatsClient) getRealTimeURL() (realtimeURL string, err error) {
	url := fmt.Sprintf("%s?appId=%s&appKey=%s", FlightStatsURL, AppID, AppKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// XXX
		return realtimeURL, err
	}
	resp, err := client.Do(req)
	if err != nil {
		// XXX
		return realtimeURL, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		response, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status: %d. resp: %s", resp.StatusCode, string(response))
	}
	var r Resp
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		// XXX
		return realtimeURL, err
	}
	realtimeURL = r.Item
	return realtimeURL, nil
}

func (c *FlightStatsClient) get() error {
	realtimeURL, err := c.getRealTimeURL()
	if err != nil {
		// XXX
		return err
	}
	url := fmt.Sprintf("%s?appId=%s&appKey=%s", realtimeURL, AppID, AppKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// XXX
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		// XXX
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		response, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("status: %d. resp: %s", resp.StatusCode, string(response))
	}
	var r FlightResp
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		// XXX
		return err
	}
	return r.Convert()
}

func StartFeeding(TotalSleepSecs float64, cargo_url string, provider string) {
	// default
	if TotalSleepSecs <= 0 {
		TotalSleepSecs = 60.0
	}
	if provider == "" {
		provider = "/HTTP_RS_SSS_Stocks"
	}

	permanently_blocked := make(chan bool)
	<-permanently_blocked
}
