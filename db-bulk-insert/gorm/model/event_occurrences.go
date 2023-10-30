package model

const TableNameEventOccurrence = "event_occurrences"

type RuleType struct {
	Startsat string `json:"startsAt"`
	Endsat   string `json:"endsAt"`
	Startson string `json:"startsOn"`
}

type EventOccurrence struct {
	// Intermediary
	Rule RuleType `gorm:"-" json:"rule"`

	// Not null
	ID int32 `gorm:"column:id;not null" json:"id"`

	// Nullable
	Startson string `gorm:"column:starts_on;default:null"`
	Startsat string `gorm:"column:starts_at;default:null"`
	Endsat   string `gorm:"column:ends_at;default:null"`

	Frequency  string `gorm:"column:frequency;default:null" json:"frequency"`
	Separation int32  `gorm:"column:separation;default:null" json:"separation"`
	Event      int32  `gorm:"column:event;default:null" json:"eventId"`
	Space      int32  `gorm:"column:space;default:null" json:"spaceId"`
}

func (EventOccurrence) TableName() string {
	return TableNameEventOccurrence
}

func (e *EventOccurrence) Format() {
	e.Startsat = e.Rule.Startsat
	e.Endsat = e.Rule.Endsat
	e.Startson = e.Rule.Startson
}
