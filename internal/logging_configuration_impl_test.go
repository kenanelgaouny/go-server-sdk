package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"gopkg.in/launchdarkly/go-sdk-common.v3/ldlog"
)

func TestLoggingConfigurationImpl(t *testing.T) {
	t.Run("GetLogDataSourceOutageAsErrorAfter", func(t *testing.T) {
		lc := LoggingConfigurationImpl{}
		assert.Equal(t, time.Duration(0), lc.GetLogDataSourceOutageAsErrorAfter())

		lc.LogDataSourceOutageAsErrorAfter = time.Second
		assert.Equal(t, time.Second, lc.GetLogDataSourceOutageAsErrorAfter())
	})

	t.Run("IsLogEvaluationErrors", func(t *testing.T) {
		lc := LoggingConfigurationImpl{}
		assert.False(t, lc.IsLogEvaluationErrors())

		lc.LogEvaluationErrors = true
		assert.True(t, lc.IsLogEvaluationErrors())
	})

	t.Run("IsLogContextKeyInErrors", func(t *testing.T) {
		lc := LoggingConfigurationImpl{}
		assert.False(t, lc.IsLogContextKeyInErrors())

		lc.LogContextKeyInErrors = true
		assert.True(t, lc.IsLogContextKeyInErrors())
	})

	t.Run("GetLoggers", func(t *testing.T) {
		loggers := ldlog.NewDefaultLoggers()
		loggers.SetMinLevel(ldlog.Error)

		lc := LoggingConfigurationImpl{}
		assert.NotEqual(t, loggers, lc.GetLoggers())

		lc.Loggers = loggers
		assert.Equal(t, loggers, lc.GetLoggers())
	})
}
