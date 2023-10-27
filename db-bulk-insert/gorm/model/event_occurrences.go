package model

import (
	"time"
)

const TableNameEventOccurrence = "event_occurrences"

// EventOccurrence mapped from table <event_occurrences>
type EventOccurrence struct {
	ID           int32     `gorm:"column:id;not null" json:"id"`
	Startson     time.Time `gorm:"column:_startson" json:"startson"`
	Endson       time.Time `gorm:"column:_endson" json:"endson"`
	Startsat     time.Time `gorm:"column:_startsat" json:"startsat"`
	Endsat       time.Time `gorm:"column:_endsat" json:"endsat"`
	Frequency    string    `gorm:"column:frequency" json:"frequency"`
	Separation   int32     `gorm:"column:separation;not null" json:"separation"`
	Count_       int32     `gorm:"column:_count" json:"count"`
	Until        time.Time `gorm:"column:_until" json:"until"`
	Timezonename string    `gorm:"column:timezonename;not null" json:"timezonename"`
	Event        int32     `gorm:"column:event" json:"event"`
	Eventid      int32     `gorm:"column:eventid;not null" json:"eventid"`
	Space        int32     `gorm:"column:space" json:"space"`
	Spaceid      int32     `gorm:"column:spaceid;not null" json:"spaceid"`
	Rule         string    `gorm:"column:rule;not null" json:"rule"`
	Status       int32     `gorm:"column:status;not null" json:"status"`
}

// TableName EventOccurrence's table name
func (*EventOccurrence) TableName() string {
	return TableNameEventOccurrence
}
