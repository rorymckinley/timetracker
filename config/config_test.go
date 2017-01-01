package config

import (
	"testing"
	"time"
)

func TestNoArgs(t *testing.T) {
	c := Config{Args: []string{}}
	conf := c.Config()
	if conf.Action != "toggle-last" {
		t.Error("Does not return toggle-last")
	}
}

func TestActionSingleArg(t *testing.T) {
	c := Config{Args: []string{"foobarbaz"}}
	conf := c.Config()
	if conf.Action != "create-new" {
		t.Error("Does not return create-new")
	}
}

func TestDescriptionSingleArg(t *testing.T) {
	c := Config{Args: []string{"foobarbaz"}, Category: "cat"}
	conf := c.Config()
	if conf.Description != "foobarbaz" {
		t.Error("Does not set description")
	}
}

func TestCategory(t *testing.T) {
	c := Config{Args: []string{"foobarbaz"}, Category: "cat"}
	conf := c.Config()
	if conf.Category != "cat" {
		t.Error("Does not set category")
	}
}

func TestCorrectlyHandlesStartTime(t *testing.T) {
	c := Config{Args: []string{"foobarbaz"}, Category: "cat", StartTime: "16:02"}
	conf := c.Config()
	location, _ := time.LoadLocation("Local")
	expected := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 2, 0, 0, location)

	if conf.StartTime != expected {
		t.Error("Does not set start time")
	}
}

func TestSetsStartTimeToNowIfNotProvided(t *testing.T) {
	format := "2006-01-02 15:04:05"
	c := Config{Args: []string{"foobarbaz"}, Category: "cat", StartTime: ""}
	conf := c.Config()
	expected := time.Now().Format(format)

	if conf.StartTime.Format(format) != expected {
		t.Error("Does not default start time")
	}
}
