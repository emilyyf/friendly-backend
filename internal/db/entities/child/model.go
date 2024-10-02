package child

import (
	"friendly-backend/internal/db/entities/household"
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/person"
	"time"

	"github.com/google/uuid"
)

type Child struct {
	ID                    uuid.UUID           `json:"id" gorm:"primaryKey"`
	IDHousehold           household.Household `json:"id_household"`
	IDMother              person.Person       `json:"id_mother"`
	Name                  string              `json:"name"`
	Birth                 time.Time           `json:"birth"`
	Age                   string              `json:"age"`
	Local                 string              `json:"local"`
	Race                  string              `json:"race"`
	AliveBirthCertificate string              `json:"alive_birth_certificate"`
	BirthCertificate      string              `json:"birth_certificate"`
	RG                    string              `json:"rg"`
	CPF                   string              `json:"cpf"`
	SUSCard               string              `json:"sus_card"`
	BloodType             string              `json:"blood_type"`
	WeightAtBirth         string              `json:"weight_at_birth"`
	HeightAtBirth         string              `json:"height_at_Birth"`
	FirstApgar            string              `json:"first_apgar"`
	FifthApgar            string              `json:"fifth_apgar"`
	NeonatalHeelPrick     time.Time           `json:"neonatal_heel_prick"`
	HearTest              time.Time           `json:"hear_test"`
	HearingTriage         time.Time           `json:"hearing_test"`
	EyeTest               string              `json:"eye_test"`
	OD                    string              `json:"od"`
	OE                    string              `json:"oe"`
	PregnancyTime         string              `json:"pregnancy_time"`
	Login                 string              `json:"login"`
	MSD                   string              `json:"msd"`
	MMII                  string              `json:"mmii"`
	CreateLog             log.Log             `json:"create_log"`
	UpdateLog             log.Log             `json:"update_log"`
}
