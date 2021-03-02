package api

import (
	fast "github.com/ddo/go-fast"
)

// NetflixSpeedTest is the implementation of SpeedTestProvider for Netfix
type NetflixSpeedTest struct {
}

// TestSpeed returns the result of the speed test or error
func (s NetflixSpeedTest) TestSpeed() (*TestResult, error) {
	fastCom := fast.New()

	// init
	err := fastCom.Init()
	if err != nil {
		return nil, err
	}

	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		return nil, err
	}

	// measure
	KbpsChan := make(chan float64)
	var downloadSpeed float64

	go func() {
		var totalMegaBytes float64
		var countUrls float64
		for Kbps := range KbpsChan {
			totalMegaBytes += Kbps/1000
			countUrls++
		}
		downloadSpeed = totalMegaBytes / countUrls
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		return nil, err
	}

	return &TestResult{Upload: 0, Download: downloadSpeed}, nil
}
