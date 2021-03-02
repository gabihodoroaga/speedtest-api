package api_test

import (
	"testing"

	"github.com/gabihodoroaga/speedtest-api/pkg/api"
	"github.com/stretchr/testify/assert"
)

// The TestSpeed shoud handle unknown provider and return a proper error
func TestUnknownProvider(t *testing.T) {
	_, err := api.TestSpeed(555)
	assert.Error(t, err)
}
