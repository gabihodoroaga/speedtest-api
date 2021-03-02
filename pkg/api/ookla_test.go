package api_test

import (
	"errors"
	"testing"

	"github.com/gabihodoroaga/speedtest-api/pkg/api/mocks"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/stretchr/testify/assert"
	"github.com/gabihodoroaga/speedtest-api/pkg/api"
)

func TestOokla(t *testing.T) {

	tests := map[string]struct {
		testLibrary api.TestLibrary
		wantErr     bool
	}{
		"FetchUserInfo returns error": {
			wantErr: true,
			testLibrary: &mockSpeedtest.TestLibraryMock{
				FetchUserInfoFunc: func() (*speedtest.User, error) {
					return nil, errors.New("Mocked fetch user error")
				},
			},
		},
		"FetchServerList returns error": {
			wantErr: true,
			testLibrary: &mockSpeedtest.TestLibraryMock{
				FetchUserInfoFunc: func() (*speedtest.User, error) {
					return &speedtest.User{}, nil
				},
				FetchServerListFunc: func(in1 *speedtest.User) (speedtest.ServerList, error) {
					return speedtest.ServerList{}, errors.New("Moked FetchServerList error")
				},
			},
		},
		"FindServer returns error": {
			wantErr: true,
			testLibrary: &mockSpeedtest.TestLibraryMock{
				FetchUserInfoFunc: func() (*speedtest.User, error) {
					return &speedtest.User{}, nil
				},
				FetchServerListFunc: func(in1 *speedtest.User) (speedtest.ServerList, error) {
					return speedtest.ServerList{}, nil
				},
				FindServerFunc: func(servers *speedtest.ServerList, serverIDs []int) (speedtest.Servers, error) {
					return nil, errors.New("Moked FindServer error")
				},
			},
		},
		"FindServer returns an empty list": {
			wantErr: true,
			testLibrary: &mockSpeedtest.TestLibraryMock{
				FetchUserInfoFunc: func() (*speedtest.User, error) {
					return &speedtest.User{}, nil
				},
				FetchServerListFunc: func(in1 *speedtest.User) (speedtest.ServerList, error) {
					return speedtest.ServerList{}, nil
				},
				FindServerFunc: func(servers *speedtest.ServerList, serverIDs []int) (speedtest.Servers, error) {
					return speedtest.Servers{}, nil
				},
			},
		},
		"DownloadTest returns error": {
			wantErr: true,
			testLibrary: &mockSpeedtest.TestLibraryMock{
				FetchUserInfoFunc: func() (*speedtest.User, error) {
					return &speedtest.User{}, nil
				},
				FetchServerListFunc: func(in1 *speedtest.User) (speedtest.ServerList, error) {
					return speedtest.ServerList{}, nil
				},
				FindServerFunc: func(servers *speedtest.ServerList, serverIDs []int) (speedtest.Servers, error) {
					targets := speedtest.Servers{}
					targets = append(targets, &speedtest.Server{})
					return targets, nil
				},
				DownloadTestFunc: func(server *speedtest.Server, savingMode bool) error {
					return errors.New("Mocked download error")
				},
			},
		},
		"UploadTest returns error": {
			wantErr: true,
			testLibrary: &mockSpeedtest.TestLibraryMock{
				FetchUserInfoFunc: func() (*speedtest.User, error) {
					return &speedtest.User{}, nil
				},
				FetchServerListFunc: func(in1 *speedtest.User) (speedtest.ServerList, error) {
					return speedtest.ServerList{}, nil
				},
				FindServerFunc: func(servers *speedtest.ServerList, serverIDs []int) (speedtest.Servers, error) {
					targets := speedtest.Servers{}
					targets = append(targets, &speedtest.Server{})
					return targets, nil
				},
				DownloadTestFunc: func(server *speedtest.Server, savingMode bool) error {
					return nil
				},
				UploadTestFunc: func(server *speedtest.Server, savingMode bool) error {
					return errors.New("Mocked upload error")
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			api.SpeedTestLibrary = tt.testLibrary
			ookla := api.OoklaSpeedTest{}
			_, err := ookla.TestSpeed()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
