package api

import (
	"errors"

	"github.com/showwin/speedtest-go/speedtest"
)

// TestLibrary is a thin wrapper for the speedtest-go so it could be mocked
//go:generate moq -pkg mockSpeedtest -out ./mocks/speedtest_mock.go . TestLibrary:TestLibraryMock
type TestLibrary interface {
	FetchUserInfo() (*speedtest.User, error)
	FetchServerList(*speedtest.User) (speedtest.ServerList, error)
	FindServer(servers *speedtest.ServerList, serverIDs []int) (
		speedtest.Servers, error)
	DownloadTest(server *speedtest.Server, savingMode bool) error
	UploadTest(server *speedtest.Server, savingMode bool) error
}

type testLibraryImpl struct {
}

func (s testLibraryImpl) FetchUserInfo() (*speedtest.User, error) {
	return speedtest.FetchUserInfo()
}

func (s testLibraryImpl) FetchServerList(user *speedtest.User) (
	speedtest.ServerList, error) {
	return speedtest.FetchServerList(user)
}

func (s testLibraryImpl) FindServer(servers *speedtest.ServerList,
	serverIDs []int) (speedtest.Servers, error) {
	return servers.FindServer(serverIDs)
}

func (s testLibraryImpl) DownloadTest(server *speedtest.Server, 
	savingMode bool) error {
	return server.DownloadTest(savingMode)
}

func (s testLibraryImpl) UploadTest(server *speedtest.Server, 
	savingMode bool) error {
		return server.UploadTest(savingMode)
}

// SpeedTestLibrary is the exposed test library so it can be mocked
var SpeedTestLibrary TestLibrary = testLibraryImpl{}

// OoklaSpeedTest is the implemenation of the SpeedTestProvider for Ookla
type OoklaSpeedTest struct {
}

// TestSpeed returns the result of the speed test or error
func (s OoklaSpeedTest) TestSpeed() (*TestResult, error) {
	user, err := SpeedTestLibrary.FetchUserInfo()
	if err != nil {
		return nil, err
	}

	serverList, err := SpeedTestLibrary.FetchServerList(user)

	if err != nil {
		return nil, err
	}

	targets, err := SpeedTestLibrary.FindServer(&serverList, []int{})
	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return nil, errors.New("no target servers found")
	}

	server := targets[0]
	if err = SpeedTestLibrary.DownloadTest(server, false); err != nil {
		return nil, err
	}

	if err = SpeedTestLibrary.UploadTest(server, false); err != nil {
		return nil, err
	}

	return &TestResult{Upload: server.ULSpeed, Download: server.DLSpeed}, nil
}
