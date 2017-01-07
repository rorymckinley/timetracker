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

func TestLastEventNoEvents(t *testing.T) {
	td := TrackingData{Events: []Event{}}

	if td.LastEvent() != nil {
		t.Error("Incorrectly returns last event on empty event set")
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

func TestCloseLastEventEmptyEvents(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Closing on Empty Event created a panic")
		}
	}()
	td := TrackingData{Events: []Event{}}
	td.CloseLastEvent()
}

func TestToggleLastLastEventOpen(t *testing.T) {
	td := TrackingData{Events: []Event{Event{}}}
	td.ToggleLast()
	if td.LastEventIsOpen() {
		t.Error("Last Event not closed by toggle")
	}
}

func TestToggleLastLastEventClosed(t *testing.T) {
	td := TrackingData{Events: []Event{Event{Description: "foo", EndTime: time.Now(), Category: "Foo",
		Subcategories: []string{"bar", "baz"}}}}
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

	if td.LastEvent().Category != closed.Category {
		t.Error("Last Event Category not set")
	}

	if !subCategoriesAreEqual(closed.Subcategories, td.LastEvent().Subcategories) {
		t.Error("Last Event Subcategories not set")
	}

	if !td.LastEventIsOpen() {
		t.Error("Last Event is not open")
	}

	if time.Since(td.LastEvent().StartTime) > time.Duration(1)*time.Second {
		t.Error("Start Time is incorrect")
	}
}

func subCategoriesAreEqual(expected []string, actual []string) bool {
	switch {
	case expected == nil && actual == nil:
		return true
	case (expected == nil && actual != nil) || (expected != nil && actual == nil):
		return false
	case len(expected) != len(actual):
		return false
	}

	for i, v := range expected {
		if actual[i] != v {
			return false
		}
	}

	return true
}
