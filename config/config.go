package config

import (
	"github.com/rorymckinley/timetracker/data"
	"strconv"
	"strings"
	"time"
)

type ActionData struct {
	Action string
	data.Data
}

type Config struct {
	Category, StartTime, EndTime, Subcategories string
	Args                                        []string
}

func (c Config) Config() ActionData {
	var config ActionData

	switch {
	case len(c.Args) == 0:
		config = ActionData{Action: "toggle-last"}
	case len(c.Args) == 1:
		data := data.Data{Description: c.Args[0], Category: c.Category, StartTime: createStart(c.StartTime)}
		config = ActionData{Action: "create-new", Data: data}
	}
	return config
}

func createStart(hoursmins string) time.Time {
	var startTime time.Time

	if hoursmins != "" {
		startTime = convert(hoursmins)
	} else {
		startTime = time.Now()
	}

	return startTime
}

func convert(hoursmins string) time.Time {
	timeparts := strings.Split(hoursmins, ":")
	now := time.Now()
	location, _ := time.LoadLocation("Local")
	hours, _ := strconv.Atoi(timeparts[0])
	minutes, _ := strconv.Atoi(timeparts[1])
	return time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, location)
	// if time != "" {
	// 	timeparts = strings.Split(time, ":")
	// 	now := Time.Now()
	// 	location, _ = time.LoadLocation("Local")
	// 	return time.Date(now.Year(), now.Date(), now.Day(), strconv.Atoi(timeparts[0]), strconv.Atoi(timeparts[1]),
	// 		0, 0, location)
	//
	// } else {
	// 	return time.Time{}
	// }

}
