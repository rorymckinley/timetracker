package data

import "time"

type Data struct {
	Description   string
	Category      string
	StartTime     time.Time
	EndTime       time.Time
	Subcategories []string
}
