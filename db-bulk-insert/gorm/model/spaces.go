package model

const TableNameSpace = "spaces"

type Space struct {
	// Intermediary
	ILocation         Location  `gorm:"-" json:"location"`
	IEventoccurrences []int32   `gorm:"-" json:"eventOccurrences"`
	IChildren         []int32   `gorm:"-" json:"children"`
	ICreatetimestamp  Timestamp `gorm:"-" json:"createTimestamp"`
	IUpdatetimestamp  Timestamp `gorm:"-" json:"updateTimestamp"`
	ITerms            TermsType `gorm:"-" json:"terms"`

	// Not null
	ID       int32  `gorm:"column:id;not null" json:"id"`
	Location string `gorm:"column:location;not null"`
	Name     string `gorm:"column:name;not null" json:"name"`

	// Nullable
	Createtimestamp  string `gorm:"column:create_timestamp;default:null"`
	Updatetimestamp  string `gorm:"column:update_timestamp;default:null"`
	Eventoccurrences string `gorm:"column:event_occurrences;default:null"`
	Children         string `gorm:"column:children;default:null"`
	Terms            string `gorm:"column:terms;default:null"`

	Shortdescription string `gorm:"column:short_description;default:null" json:"shortDescription"`
	Horario          string `gorm:"column:horario;default:null" json:"horario"`
	Telefone         string `gorm:"column:telefone;default:null" json:"telefonePublico"`
	Email            string `gorm:"column:email;default:null" json:"emailPublico"`
	Parent           int32  `gorm:"column:parent;default:null" json:"parent"`
	Owner            int32  `gorm:"column:owner;default:null" json:"owner"`
}

func (Space) TableName() string {
	return TableNameSpace
}

func (s *Space) Format() {
	s.Eventoccurrences = Format(s.IEventoccurrences)
	s.Children = Format(s.IChildren)
	s.Location = s.ILocation.Latitude + "," + s.ILocation.Longitude

	if s.Location == "0,0" {
		s.Location = ""
	}

	s.Createtimestamp = FormatTimestamp(s.ICreatetimestamp)
	s.Updatetimestamp = FormatTimestamp(s.IUpdatetimestamp)
	s.Terms = FormatQuery(s.ITerms)
}
