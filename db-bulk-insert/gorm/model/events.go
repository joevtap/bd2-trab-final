package model

import (
	"time"
)

const TableNameEvent = "events"

// Event mapped from table <events>
type Event struct {
	ID               int32     `gorm:"column:id;not null" json:"id"`
	Type             int32     `gorm:"column:_type;not null" json:"type"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Shortdescription string    `gorm:"column:shortdescription;not null" json:"shortdescription"`
	Longdescription  string    `gorm:"column:longdescription" json:"longdescription"`
	Rules            string    `gorm:"column:rules" json:"rules"`
	Createtimestamp  time.Time `gorm:"column:createtimestamp" json:"createtimestamp"`
	Status           int32     `gorm:"column:status;not null" json:"status"`
	IOccurrences     []int32   `gorm:"-" json:"occurrences"`
	Occurrences      string    `gorm:"column:occurrences"`
	Owner            int32     `gorm:"column:owner" json:"owner"`
	Project          int32     `gorm:"column:project" json:"project"`
}

// TableName Event's table name
func (*Event) TableName() string {
	return TableNameEvent
}

func (e *Event) Format() {
	e.Occurrences = Format(e.IOccurrences)
}
