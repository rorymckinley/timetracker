package trackingdata

import (
	"testing"
)

func TestLastEvent(t *testing.T) {
	td := TrackingData{Events: []Event{Event{}, Event{}}}
	if td.LastEvent() != &td.Events[1] {
		t.Error("Does not return last event")
	}
}
