package model

const TableNameAgent = "agents"

type Agent struct {
	// Intermediary
	ICreatetimestamp Timestamp `gorm:"-" json:"createTimestamp"`
	IUpdatetimestamp Timestamp `gorm:"-" json:"updateTimestamp"`
	IChildren        []int32   `gorm:"-" json:"children"`
	ISpaces          []int32   `gorm:"-" json:"spaces"`
	IProjects        []int32   `gorm:"-" json:"projects"`
	IEvents          []int32   `gorm:"-" json:"events"`
	ITerms           TermsType `gorm:"-" json:"terms"`

	// Not null
	ID   int32  `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`

	// Nullable
	Createtimestamp string `gorm:"column:create_timestamp;default:null"`
	Updatetimestamp string `gorm:"column:update_timestamp;default:null"`
	Children        string `gorm:"column:children;default:null"`
	Spaces          string `gorm:"column:spaces;default:null"`
	Projects        string `gorm:"column:projects;default:null"`
	Events          string `gorm:"column:events;default:null"`
	Terms           string `gorm:"terms;default:null"`

	Shortdescription string `gorm:"column:short_description;default:null" json:"shortDescription"`
	Parent           int32  `gorm:"column:parent;default:null" json:"parent"`
}

func (Agent) TableName() string {
	return TableNameAgent
}

func (a *Agent) Format() {
	a.Children = Format(a.IChildren)
	a.Spaces = Format(a.ISpaces)
	a.Projects = Format(a.IProjects)
	a.Events = Format(a.IEvents)
	a.Createtimestamp = FormatTimestamp(a.ICreatetimestamp)
	a.Updatetimestamp = FormatTimestamp(a.IUpdatetimestamp)
	a.Terms = FormatQuery(a.ITerms)
}
