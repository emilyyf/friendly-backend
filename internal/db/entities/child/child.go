package child

import (
	"friendly-backend/internal/db/entities/household"
	"friendly-backend/internal/db/entities/log"
	"friendly-backend/internal/db/entities/person"
	"time"

	"github.com/google/uuid"
)

type Child struct {
	ID                    uuid.UUID           `json:"id"`
	IDHousehold           household.Household `json:"id_household"`
	IDMother              person.Person       `json:"id_mother"`
	Name                  string              `json:"name"`
	Birth                 time.Time           `json:"birth"`
	Age                   int64               `json:"age"`
	Local                 string              `json:"local"`
	Race                  string              `json:"race"`
	AliveBirthCertificate string              `json:"alive_birth_certificate"`
	BirthCertificate      string              `json:"birth_certificate"`
	RG                    string              `json:"rg"`
	CPF                   string              `json:"cpf"`
	SUSCard               int64               `json:"sus_card"`
	BloodType             string              `json:"blood_type"`
	WeightAtBirth         int64               `json:"weight_at_birth"`
	HeightAtBirth         int64               `json:"height_at_Birth"`
	FirstApgar            int64               `json:"first_apgar"`
	FifthApgar            int64               `json:"fifth_apgar"`
	NeonatalHeelPrick     int64               `json:"neonatal_heel_prick"`
	HearTest              int64               `json:"hear_test"`
	HearingTriage         int64               `json:"hearing_test"`
	EyeTest               int64               `json:"eye_test"`
	OD                    int64               `json:"od"`
	OE                    int64               `json:"oe"`
	PregnancyTime         int64               `jsno:"pregnancy_time"`
	Login                 int64               `json:"login"`
	MSD                   int64               `json:"msd"`
	MMII                  int64               `json:"mmii"`
	CreateLog             log.Log             `json:"create_log"`
	UpdateLog             log.Log             `json:"update_log"`
}
