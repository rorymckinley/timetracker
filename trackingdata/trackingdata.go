package trackingdata

import (
	"github.com/rorymckinley/timetracker/data"
	"time"
)

type Event struct {
	Description   string    `yaml: "description"`
	StartTime     time.Time `yaml: "startTime"`
	EndTime       time.Time `yaml: "endTime"`
	Category      string    `yaml: category`
	Subcategories []string
}

type TrackingData struct {
	Events []Event
}

func (td *TrackingData) LastEvent() *Event {
	if len(td.Events) > 0 {
		return &td.Events[len(td.Events)-1]
	} else {
		return nil
	}
}

func (td *TrackingData) LastEventIsOpen() bool {
	return td.LastEvent().EndTime.IsZero()
}

func (td *TrackingData) CloseLastEvent() {
	if last := td.LastEvent(); last != nil {
		last.EndTime = time.Now()
	}
}

func (td *TrackingData) ToggleLast() {
	if td.LastEventIsOpen() {
		td.CloseLastEvent()
	} else {
		td.AddEvent(data.Data{StartTime: time.Now(), Description: td.LastEvent().Description,
			Category: td.LastEvent().Category, Subcategories: td.LastEvent().Subcategories})
	}
}

func (td *TrackingData) AddEvent(data data.Data) {
	events := append(td.Events, startEvent(data))
	td.Events = events
}

func startEvent(data data.Data) (event Event) {
	return Event{Description: data.Description, Category: data.Category, StartTime: data.StartTime, EndTime: data.EndTime, Subcategories: data.Subcategories}
}
