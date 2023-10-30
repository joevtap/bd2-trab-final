package model

import "fmt"

const TableNameEvent = "events"

type Event struct {
	// Intermediary
	Terms            TermsType `gorm:"-" json:"terms"`
	IUpdatetimestamp Timestamp `gorm:"-" json:"updateTimestamp"`
	ICreatetimestamp Timestamp `gorm:"-" json:"createTimestamp"`

	// Not null
	ID   int32  `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`

	// Nullable
	Shortdescription    string `gorm:"column:short_description;default:null" json:"shortDescription"`
	Createtimestamp     string `gorm:"column:create_timestamp;default:null"`
	Updatetimestamp     string `gorm:"column:update_timestamp;default:null"`
	Owner               int32  `gorm:"column:owner;default:null" json:"owner"`
	Project             int32  `gorm:"column:project;default:null" json:"project"`
	Q                   string `gorm:"column:q;default:null"`
	ClassificacaoEtaria string `gorm:"column:classificacao_etaria;default:null" json:"classificacaoEtaria"`
}

func (Event) TableName() string {
	return TableNameEvent
}

func (e *Event) Format() {
	e.Createtimestamp = FormatTimestamp(e.ICreatetimestamp)
	e.Updatetimestamp = FormatTimestamp(e.IUpdatetimestamp)
	e.Q = fmt.Sprintf("%v %v %v %v", e.Name, e.Shortdescription, e.ClassificacaoEtaria, FormatQuery(e.Terms))
}
