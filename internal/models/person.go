package models

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID              uuid.UUID `json:"id"`
	IDHousehould    Household `json:"id_househould"`
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
	CarUF           string    `json:"card_uf"`
	Company         string    `json:"company"`
	WorkFunction    string    `json:"work_function"`
	EnploymentCard  int64     `json:"enployment_card"`
	HiringDate      time.Time `json:"hiring_date"`
	ResignationDate time.Time `json:"resignation_date"`
	Salary          int64     `json:"salary"`
	ExtraIncome     int64     `json:"extra_income"`
	CreateLog       Log       `json:"create_log"`
	UpdateLog       Log       `json:"update_log"`
}
