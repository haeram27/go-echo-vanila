package restapis

import (
	"context"
	"echoinit/apps"
	"io/ioutil"
	"net/http"
	"time"
)

func CoronaStatus(ctx context.Context) error {
	key := ""
	api := `http://apis.data.go.kr/1790387/covid19CurrentStatusKorea/covid19CurrentStatusKoreaJason?serviceKey=`
	req, err := http.NewRequest("GET", api+key, nil)
	if err != nil {
		apps.Logs.Error(err)
	}

	// Header
	//req.Header.Add("User-Agent", "Crawler")

	client := &http.Client{}
	client.Timeout = time.Second * 3

	startTime := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		apps.Logs.Error(err)
	}
	defer resp.Body.Close()
	elapsedTime := time.Since(startTime)
	apps.Logs.Debug("elapsed:", elapsedTime)
	bytes, _ := ioutil.ReadAll(resp.Body)
	apps.Logs.Info(string(bytes))

	return nil
}
