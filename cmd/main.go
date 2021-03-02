package main

import (
	"log"

	"github.com/gabihodoroaga/speedtest-api/pkg/api"
)

func main() {
	log.Printf("Begin speed test\n")
	testResult, err := api.TestSpeed(api.Ookla)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Ookla speed test - download: %.2f Mbps, upload: %.2f Mbps", 
		testResult.Download, testResult.Upload)

	testResult, err = api.TestSpeed(api.Netflix)
	log.Printf("Netflix speed test - download: %.2f Mbps, upload: %.2f Mbps", 
		testResult.Download, testResult.Upload)

	log.Printf("Done")
}
