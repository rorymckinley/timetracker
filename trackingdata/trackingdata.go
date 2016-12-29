package trackingdata

import (
	"time"
)

type Event struct {
	Description   string    `yaml: "description"`
	StartTime     time.Time `yaml: "startTime"`
	EndTime       time.Time `yaml: "endTime"`
	Category      string    `yaml: category`
	SubCategories []string
}

type TrackingData struct {
	Events []Event
}

func (td *TrackingData) LastEvent() *Event {
	return &td.Events[len(td.Events)-1]
}
