package loglevel

import (
	"testing"

	"github.com/shellhub-io/shellhub/pkg/envs"
	envMocks "github.com/shellhub-io/shellhub/pkg/envs/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLevels(t *testing.T) {
	mocks := &envMocks.Backend{}
	envs.DefaultBackend = mocks

	cases := []struct {
		description   string
		requiredMocks func()
		expected      logrus.Level
	}{
		{
			description: "Set to info when variable SHELLHUB_ENV value is empty",
			requiredMocks: func() {
				mocks.On("Get", "SHELLHUB_ENV").Return("").Once()
				mocks.On("Get", "SHELLHUB_LOG_LEVEL").Return("").Once()
			},
			expected: logrus.InfoLevel,
		}, {
			description: "Set loglevel to trace when SHELLHUB_ENV is set to development",
			requiredMocks: func() {
				mocks.On("Get", "SHELLHUB_ENV").Return("development").Once()
				mocks.On("Get", "SHELLHUB_LOG_LEVEL").Return("").Once()
			},
			expected: logrus.TraceLevel,
		}, {
			description: "Set loglevel to debug when SHELLHUB_LOG_LEVEL is set to debug ",
			requiredMocks: func() {
				mocks.On("Get", "SHELLHUB_ENV").Return("").Once()
				mocks.On("Get", "SHELLHUB_LOG_LEVEL").Return("debug").Once()
			},
			expected: logrus.DebugLevel,
		}, {
			description: "Set loglevel to error when SHELLHUB_LOG_LEVEL is set to error ",
			requiredMocks: func() {
				mocks.On("Get", "SHELLHUB_ENV").Return("").Once()
				mocks.On("Get", "SHELLHUB_LOG_LEVEL").Return("error").Once()
			},
			expected: logrus.ErrorLevel,
		},
	}
	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			tc.requiredMocks()

			SetLogLevel()
			assert.Equal(t, tc.expected, logrus.GetLevel())
		})
	}
}
