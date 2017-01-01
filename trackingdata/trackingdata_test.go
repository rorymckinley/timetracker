package trackingdata

import (
	"testing"
	"time"
)

func TestLastEvent(t *testing.T) {
	td := TrackingData{Events: []Event{Event{}, Event{}}}
	if td.LastEvent() != &td.Events[1] {
		t.Error("Does not return last event")
	}
}

func TestLastEventIsOpen(t *testing.T) {
	td := TrackingData{Events: []Event{Event{}}}
	if !td.LastEventIsOpen() {
		t.Error("Does not report on last event open")
	}
}

func TestLastEventIsClosed(t *testing.T) {
	td := TrackingData{Events: []Event{Event{EndTime: time.Now()}}}
	if td.LastEventIsOpen() {
		t.Error("Does not report on last event open")
	}
}

func TestClosesLastEvent(t *testing.T) {
	td := TrackingData{Events: []Event{Event{}}}
	td.CloseLastEvent()
	if td.LastEventIsOpen() {
		t.Error("Last Event not closed")
	}
}

func TestToggleLastLastEventOpen(t *testing.T) {
	td := TrackingData{Events: []Event{Event{}}}
	td.ToggleLast()
	if td.LastEventIsOpen() {
		t.Error("Last Event not closed by toggle")
	}
}

func TestToggleLastLastEventClosed(t *testing.T) {
	td := TrackingData{Events: []Event{Event{Description: "foo", EndTime: time.Now()}}}
	closed := &td.Events[0]
	td.ToggleLast()

	if len(td.Events) != 2 {
		t.Error("Incorrect number of events")
	}
	if td.LastEvent() == closed {
		t.Error("Incorrect last event")
	}
	if td.LastEvent().Description != closed.Description {
		t.Error("Last Event Description not set")
	}
	if !td.LastEventIsOpen() {
		t.Error("Last Event is not open")
	}
	if time.Since(td.LastEvent().StartTime) > time.Duration(1)*time.Second {
		t.Error("Start Time is incorrect")
	}
}
