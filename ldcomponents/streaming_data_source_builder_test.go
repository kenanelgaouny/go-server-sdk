package ldcomponents

import (
	"testing"
	"time"

	"gopkg.in/launchdarkly/go-server-sdk.v6/internal/datasource"
	"gopkg.in/launchdarkly/go-server-sdk.v6/internal/datastore"
	"gopkg.in/launchdarkly/go-server-sdk.v6/internal/sharedtest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamingDataSourceBuilder(t *testing.T) {
	t.Run("InitialReconnectDelay", func(t *testing.T) {
		s := StreamingDataSource()
		assert.Equal(t, DefaultInitialReconnectDelay, s.initialReconnectDelay)

		s.InitialReconnectDelay(time.Minute)
		assert.Equal(t, time.Minute, s.initialReconnectDelay)

		s.InitialReconnectDelay(0)
		assert.Equal(t, DefaultInitialReconnectDelay, s.initialReconnectDelay)

		s.InitialReconnectDelay(-1 * time.Millisecond)
		assert.Equal(t, DefaultInitialReconnectDelay, s.initialReconnectDelay)
	})

	t.Run("CreateDataSource", func(t *testing.T) {
		baseURI := "base"
		delay := time.Hour

		s := StreamingDataSource().InitialReconnectDelay(delay)

		dsu := sharedtest.NewMockDataSourceUpdates(datastore.NewInMemoryDataStore(sharedtest.NewTestLoggers()))
		ds, err := s.CreateDataSource(makeTestContextWithBaseURIs(baseURI), dsu)
		require.NoError(t, err)
		require.NotNil(t, ds)
		defer ds.Close()

		sp := ds.(*datasource.StreamProcessor)
		assert.Equal(t, baseURI, sp.GetBaseURI())
		assert.Equal(t, delay, sp.GetInitialReconnectDelay())
	})
}
