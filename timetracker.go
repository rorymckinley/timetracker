package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var Category = flag.String("cat", "", "Category")
var StartTime = flag.String("start", "", "Start Time")
var EndTime = flag.String("end", "", "End Time")
var Subs = flag.String("subs", "", "Subcategories")

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
	events := append(td.Events, startEvent(data))
	td.Events = events
}

type Data struct {
	Description   string
	Category      string
	StartTime     time.Time
	EndTime       time.Time
	SubCategories []string
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
		SetCategory(&data)
		SetStartTime(&data)
		SetEndTime(&data)
		SetSubCategories(&data)
		return ActionData{Action: "create-new", Data: data}
	default:
		return ActionData{}
	}
}

func SetCategory(data *Data) {
	if *Category != "" {
		data.Category = *Category
	} else {
		fmt.Println("Please supply a category")
		os.Exit(1)
	}
}

func SetStartTime(data *Data) {
	const dateTemplate = "2006 January 2"
	const timeTemplate = "2006 January 2 15:04 -0700"
	if *StartTime != "" {
		input := time.Now().Format(dateTemplate) + " " + *StartTime + " +0200"
		data.StartTime, _ = time.Parse(timeTemplate, input)
	} else {
		data.StartTime = time.Now()
	}
}

func SetEndTime(data *Data) {
	const dateTemplate = "2006 January 2"
	const timeTemplate = "2006 January 2 15:04 -0700"
	if *EndTime != "" {
		input := time.Now().Format(dateTemplate) + " " + *EndTime + " +0200"
		data.EndTime, _ = time.Parse(timeTemplate, input)
	}
}

func SetSubCategories(data *Data) {
	data.SubCategories = strings.Split(*Subs, ",")
}

func startEvent(data Data) (event Event) {
	return Event{Description: data.Description, Category: data.Category, StartTime: data.StartTime, EndTime: data.EndTime, SubCategories: data.SubCategories}
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
