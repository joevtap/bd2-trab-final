package model

const TableNameProject = "projects"

type Project struct {
	// Intermediary
	IEvents           []int32   `gorm:"-" json:"events"`
	IChildren         []int32   `gorm:"-" json:"children"`
	ICreatetimestamp  Timestamp `gorm:"-" json:"createTimestamp"`
	IUpdatetimestamp  Timestamp `gorm:"-" json:"updateTimestamp"`
	IRegistrationfrom Timestamp `gorm:"-" json:"registrationFrom"`
	IRegistrationto   Timestamp `gorm:"-" json:"registrationTo"`

	// Not null
	ID   int32  `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`

	// Nullable
	Shortdescription string `gorm:"column:short_description;default:null" json:"shortDescription"`
	Createtimestamp  string `gorm:"column:create_timestamp;default:null"`
	Updatetimestamp  string `gorm:"column:update_timestamp;default:null"`
	Registrationfrom string `gorm:"column:registration_from;default:null"`
	Registrationto   string `gorm:"column:registration_to;default:null"`

	Children string `gorm:"column:children;default:null"`
	Events   string `gorm:"column:events;default:null"`

	Parent int32 `gorm:"column:parent;default:null" json:"parent"`
	Owner  int32 `gorm:"column:owner;default:null" json:"owner"`
}

func (Project) TableName() string {
	return TableNameProject
}

func (p *Project) Format() {
	p.Children = Format(p.IChildren)
	p.Events = Format(p.IEvents)
	p.Updatetimestamp = FormatTimestamp(p.IUpdatetimestamp)
	p.Registrationfrom = FormatTimestamp(p.IRegistrationfrom)
	p.Registrationto = FormatTimestamp(p.IRegistrationto)
	p.Createtimestamp = FormatTimestamp(p.ICreatetimestamp)
}
