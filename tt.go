package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type Event struct {
	Description string    `yaml: "description"`
	StartTime   time.Time `yaml: "startTime"`
	EndTime     time.Time `yaml: "endTime"`
}

type TrackingData struct {
	Events []Event
}

func (td *TrackingData) LastEvent() *Event {
	return &td.Events[len(td.Events)-1]
}

func (td *TrackingData) LastEventIsOpen() bool {
	return td.LastEvent().EndTime.IsZero()
}

func (td *TrackingData) CloseLastEvent() {
	td.LastEvent().EndTime = time.Now()
}

func (td *TrackingData) ToggleLast() {
	if td.LastEventIsOpen() {
		td.CloseLastEvent()
	} else {
		td.AddEvent(Data{Description: td.LastEvent().Description})
	}
}

func (td *TrackingData) AddEvent(data Data) {
	events := append(td.Events, startEvent(data.Description))
	td.Events = events
}

type Data struct {
	Description string
}

type ActionData struct {
	Action string
	Data
}

func read(location string) TrackingData {
	trackingData := TrackingData{}

	fileData, _ := ioutil.ReadFile(location)
	err := yaml.Unmarshal(fileData, &trackingData)

	if err != nil {
		spew.Dump(err)
	}

	return trackingData
}

func persist(data *TrackingData, location string) {
	export, _ := yaml.Marshal(data)
	f, _ := os.Create(location)
	f.Write(export)
	f.Close()
}

func determineConfig() ActionData {
	switch {
	case len(flag.Args()) == 0:
		return ActionData{Action: "toggle-last"}
	case len(flag.Args()) == 1:
		data := Data{Description: flag.Arg(0)}
		return ActionData{Action: "create-new", Data: data}
	default:
		return ActionData{}
	}
}

func startEvent(description string) (event Event) {
	return Event{Description: description, StartTime: time.Now()}
}

// Called without args: If running - stop current running event - if stopped, create a copy of last event and start it
func main() {
	flag.Parse()
	config := determineConfig()

	trackingData := read("./tracking.yml")

	switch config.Action {
	case "toggle-last":
		trackingData.ToggleLast()
	case "create-new":
		trackingData.CloseLastEvent()
		trackingData.AddEvent(config.Data)
	}

	persist(&trackingData, "./tracking.yml")
}
