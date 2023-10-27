package model

const TableNameSpace = "spaces"

// Space mapped from table <spaces>
type Space struct {
	ID                int32   `gorm:"column:id;not null" json:"id"`
	Location          string  `gorm:"column:location;not null" json:"location"`
	Geolocation       string  `gorm:"column:_geolocation;not null" json:"geolocation"`
	Name              string  `gorm:"column:name;not null" json:"name"`
	Public            bool    `gorm:"column:public;not null" json:"public"`
	Shortdescription  string  `gorm:"column:shortdescription" json:"shortdescription"`
	Longdescription   string  `gorm:"column:longdescription" json:"longdescription"`
	Status            int32   `gorm:"column:status;not null" json:"status"`
	Type              int32   `gorm:"column:_type;not null" json:"type"`
	IEventoccurrences []int32 `gorm:"-" json:"eventOccurrences"`
	Eventoccurrences  string  `gorm:"column:eventoccurrences"`
	Parent            int32   `gorm:"column:parent" json:"parent"`
	IChildren         []int32 `gorm:"-" json:"children"`
	Children          string  `gorm:"column:_children"`
	Owner             int32   `gorm:"column:owner" json:"owner"`
}

// TableName Space's table name
func (*Space) TableName() string {
	return TableNameSpace
}

func (s *Space) Format() {
	s.Eventoccurrences = Format(s.IEventoccurrences)
	s.Children = Format(s.IChildren)
}
