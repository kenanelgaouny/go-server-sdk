package ldclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/launchdarkly/go-sdk-common.v2/ldvalue"
)

func TestFlagsStateCanGetFlagValue(t *testing.T) {
	flag := FeatureFlag{Key: "key"}
	state := newFeatureFlagsState()
	state.addFlag(&flag, ldvalue.String("value"), intPtr(1), EvaluationReason{}, false)

	assert.Equal(t, ldvalue.String("value"), state.GetFlagValue("key"))
}

func TestFlagsStateUnknownFlagReturnsNilValue(t *testing.T) {
	state := newFeatureFlagsState()

	assert.Equal(t, ldvalue.Null(), state.GetFlagValue("key"))
}

func TestFlagsStateCanGetFlagReason(t *testing.T) {
	flag := FeatureFlag{Key: "key"}
	state := newFeatureFlagsState()
	state.addFlag(&flag, ldvalue.String("value"), intPtr(1), newEvalReasonOff(), false)

	assert.Equal(t, newEvalReasonOff(), state.GetFlagReason("key"))
}

func TestFlagsStateUnknownFlagReturnsEmptyReason(t *testing.T) {
	state := newFeatureFlagsState()

	assert.Equal(t, EvaluationReason{}, state.GetFlagReason("key"))
}

func TestFlagsStateReturnsEmptyReasonIfReasonsWereNotRecorded(t *testing.T) {
	flag := FeatureFlag{Key: "key"}
	state := newFeatureFlagsState()
	state.addFlag(&flag, ldvalue.String("value"), intPtr(1), EvaluationReason{}, false)

	assert.Equal(t, EvaluationReason{}, state.GetFlagReason("key"))
}

func TestFlagsStateToValuesMap(t *testing.T) {
	flag1 := FeatureFlag{Key: "key1"}
	flag2 := FeatureFlag{Key: "key2"}
	state := newFeatureFlagsState()
	state.addFlag(&flag1, ldvalue.String("value1"), intPtr(0), EvaluationReason{}, false)
	state.addFlag(&flag2, ldvalue.String("value2"), intPtr(1), EvaluationReason{}, false)

	expected := map[string]ldvalue.Value{"key1": ldvalue.String("value1"), "key2": ldvalue.String("value2")}
	assert.Equal(t, expected, state.ToValuesMap())
}

func TestFlagsStateToJSON(t *testing.T) {
	date := uint64(1000)
	flag1 := FeatureFlag{Key: "key1", Version: 100, TrackEvents: false}
	flag2 := FeatureFlag{Key: "key2", Version: 200, TrackEvents: true, DebugEventsUntilDate: &date}
	state := newFeatureFlagsState()
	state.addFlag(&flag1, ldvalue.String("value1"), intPtr(0), EvaluationReason{}, false)
	state.addFlag(&flag2, ldvalue.String("value2"), intPtr(1), EvaluationReason{}, false)

	expectedString := `{
		"key1":"value1",
		"key2":"value2",
		"$flagsState":{
	  		"key1":{
				"variation":0,"version":100,"reason":null
			},
			"key2": {
				"variation":1,"version":200,"trackEvents":true,"debugEventsUntilDate":1000,"reason":null
			}
		},
		"$valid":true
	}`
	actualBytes, err := json.Marshal(state)
	assert.NoError(t, err)
	assert.JSONEq(t, expectedString, string(actualBytes))
}
