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

func determineConfig() map[string]string {
	if len(flag.Args()) == 0 {
		return map[string]string{"action": "toggle-last"}
	}
	return map[string]string{}
}

func toggleLast(data *TrackingData) {
	lastEvent := &data.Events[len(data.Events)-1]
	if lastEvent.EndTime.IsZero() {
		lastEvent.EndTime = time.Now()
	} else {
		data.Events = append(data.Events, startEvent(lastEvent.Description))
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

	switch config["action"] {
	case "toggle-last":
		toggleLast(&trackingData)
	}
	// description := os.Args[1]
	//
	// newEvent :=
	//
	// trackingData.Events = append(trackingData.Events, newEvent)

	persist(&trackingData, "./tracking.yml")
}
