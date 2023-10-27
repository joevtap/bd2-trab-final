package model

import (
	"time"
)

const TableNameProject = "projects"

// Project mapped from table <projects>
type Project struct {
	ID               int32     `gorm:"column:id;not null" json:"id"`
	Type             int32     `gorm:"column:_type;not null" json:"type"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Shortdescription string    `gorm:"column:shortdescription" json:"shortdescription"`
	Longdescription  string    `gorm:"column:longdescription" json:"longdescription"`
	Updatetimestamp  time.Time `gorm:"column:updatetimestamp" json:"updatetimestamp"`
	Registrationfrom time.Time `gorm:"column:registrationfrom" json:"registrationfrom"`
	Registrationto   time.Time `gorm:"column:registrationto" json:"registrationto"`
	Createtimestamp  time.Time `gorm:"column:createtimestamp" json:"createtimestamp"`
	Status           int32     `gorm:"column:status;not null" json:"status"`
	Parent           int32     `gorm:"column:parent" json:"parent"`
	IChildren        []int32   `gorm:"-" json:"children"`
	Children         string    `gorm:"column:_children"`
	Owner            int32     `gorm:"column:owner" json:"owner"`
	IEvents          []int32   `gorm:"-" json:"events"`
	Events           string    `gorm:"column:_events" `
}

// TableName Project's table name
func (*Project) TableName() string {
	return TableNameProject
}

func (p *Project) Format() {
	p.Children = Format(p.IChildren)
	p.Events = Format(p.IEvents)
}
