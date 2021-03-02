# speedtest-api

A very simple library that tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.

## The solution

The solution is rather complicated then simple because the purpose of this library is to showcase how you can create different implementations for an interface and how you can mock external packages by creating wrapping interfaces and generate mocks using the moq tool [github.com/matryer/moq](https://github.com/matryer/moq).

## Known issues

- Due to time limitation I used 2 external libraries to test the download speed from Ookla and Neflix and these libraries are not testable and also not mockable
- No benchmark tests. A benchmark tests the performance of a function by calling the function repetitively. This library forwards the requests to the underlying libraries which contains http calls
- No test for the Netflix implementation

## How to use

```bash
go get github.com/gabihodoroaga/speedtest-api
```

```go
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
```

## TODO:

- rewrite the speed testing library so can be testable, mockable and allow canceling of the speed test
