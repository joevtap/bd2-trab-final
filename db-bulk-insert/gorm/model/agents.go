package model

const TableNameAgent = "agents"

// Agent mapped from table <agents>
type Agent struct {
	ID               int32   `gorm:"column:id;not null" json:"id"`
	Type             int32   `gorm:"column:_type;not null" json:"type"`
	Name             string  `gorm:"column:name;not null" json:"name"`
	Shortdescription string  `gorm:"column:shortdescription" json:"shortdescription"`
	Longdescription  string  `gorm:"column:longdescription" json:"longdescription"`
	Status           int32   `gorm:"column:status;not null" json:"status"`
	Parent           int32   `gorm:"column:parent" json:"parent"`
	IChildren        []int32 `gorm:"-" json:"children"`
	ISpaces          []int32 `gorm:"-" json:"spaces"`
	IProjects        []int32 `gorm:"-" json:"projects"`
	IEvents          []int32 `gorm:"-" json:"events"`
	Children         string  `gorm:"column:_children"`
	Spaces           string  `gorm:"column:_spaces"`
	Projects         string  `gorm:"column:_projects"`
	Events           string  `gorm:"column:_events"`
}

// TableName Agent's table name
func (*Agent) TableName() string {
	return TableNameAgent
}

func (a *Agent) Format() {
	a.Children = Format(a.IChildren)
	a.Spaces = Format(a.ISpaces)
	a.Projects = Format(a.IProjects)
	a.Events = Format(a.IEvents)
}
