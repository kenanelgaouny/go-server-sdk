package ldcomponents

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/launchdarkly/go-sdk-common/v3/lduser"
	ldevents "github.com/launchdarkly/go-sdk-events/v2"
)

func TestNoEvents(t *testing.T) {
	ep, err := NoEvents().Build(basicClientContext())
	require.NoError(t, err)
	defer ep.Close()
	ef := ldevents.NewEventFactory(false, nil)
	ep.RecordIdentifyEvent(ef.NewIdentifyEventData(ldevents.Context(lduser.NewUser("key"))))
	ep.Flush()
}
