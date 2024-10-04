package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	HouseholdID     uuid.UUID `json:"id_household" gorm:"type:uuid;"`
	Household       Household
	Birth           time.Time `json:"birth"`
	Age             string    `json:"age"`
	Name            string    `json:"name"`
	Country         string    `json:"country"`
	UF              string    `json:"uf"`
	Phone           string    `json:"phone"`
	Graduation      string    `json:"graduation"`
	RG              string    `json:"rg"`
	RGEXP           time.Time `json:"rg_exp"`
	CPF             string    `json:"cpf"`
	SusCard         string    `json:"sus_card"`
	CardSeries      string    `json:"card_series"`
	CardUF          string    `json:"card_uf"`
	Company         string    `json:"company"`
	WorkFunction    string    `json:"work_function"`
	EnploymentCard  string    `json:"enployment_card"`
	HiringDate      time.Time `json:"hiring_date"`
	ResignationDate time.Time `json:"resignation_date"`
	Salary          int64     `json:"salary"`
	ExtraIncome     int64     `json:"extra_income"`
	CreateLogID     uuid.UUID `json:"create_log" gorm:"type:uuid;"`
	CreateLog       Log
	UpdateLogID     uuid.UUID `json:"update_log" gorm:"type:uuid;"`
	UpdateLog       Log
}

func (p *Person) BeforeCreate(d *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
