package household

import (
	"friendly-backend/internal/db/entities/log"
	"time"

	"github.com/google/uuid"
)

type Household struct {
	ID                uuid.UUID `json:"id"`
	Code              int64     `json:"adress"`
	Date              time.Time `json:"date"`
	Adress            string    `json:"adress"`
	AdressNumber      string    `json:"adress_number"`
	AdressComplement  string    `json:"adress_complement"`
	CEP               string    `json:"cep"`
	City              string    `json:"city"`
	Neighborhood      string    `json:"neighborhood"`
	ResidenseType     string    `json:"residense_type"`
	RentValue         int64     `json:"rent_value"`
	BuildingMaterials string    `json:"building_materials"`
	InHouseBathroom   bool      `json:"in_house_bathroom"`
	Residents         int64     `json:"residents"`
	Rooms             int64     `json:"rooms"`
	Beds              int64     `json:"beds"`
	Car               bool      `json:"car"`
	Television        bool      `json:"television"`
	Refrigerator      bool      `json:"refrigerator"`
	Microwave         bool      `json:"microwave"`
	WashingMachine    bool      `json:"washing_machine"`
	RoadType          string    `json:"road_type"`
	ReferingPerson    string    `json:"refering_person"`
	CreateLog         log.Log   `json:"create_log"`
	UpdateLog         log.Log   `json:"update_log"`
}
