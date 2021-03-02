package api

import (
	"fmt"
)

// TestResult is the result of a speed test
type TestResult struct {
	// The upload speed in Mbps
	Upload float64
	// The download speed in Mbps
	Download float64
}

// TestProvider is the type of the test provider
type TestProvider int

const (
	// Ookla is the Ookla test provider using the speedtest.net
	Ookla TestProvider = iota
	// Netflix is the Netflix test provider using the fast.com
	Netflix
)

type speedTestProvider interface {
	TestSpeed() (*TestResult, error)
}

// TestSpeed returns the result of the test using the specified test provider
// or error
func TestSpeed(provider TestProvider) (*TestResult, error) {

	var testProvider speedTestProvider
	switch provider {
	case Ookla:
		testProvider = OoklaSpeedTest{}
	case Netflix:
		testProvider = NetflixSpeedTest{}
	default:
		return nil, fmt.Errorf("unknown test provider %d", provider)
	}

	return testProvider.TestSpeed()
}
